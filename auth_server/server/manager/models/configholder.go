package models

// TODO it is too large
import (
	"fmt"
	"github.com/cesanta/docker_auth/auth_server/authn"
	"github.com/cesanta/docker_auth/auth_server/authz"
	"github.com/cesanta/docker_auth/auth_server/server/config"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type ByString [][]string

func (s ByString) Len() int {
	return len(s)
}

func (s ByString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByString) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

type AuthConfig struct {
	Config         *config.Config
	Authenticators []authn.Authenticator
	Authorizers    []authz.Authorizer
}

type AuthConfigManager struct {
	authConfig  *AuthConfig
	session     *mgo.Session
	userSession *mgo.Session
	lock        sync.Mutex
}

var ACManager *AuthConfigManager

func InitAuthConfigManager(config *config.Config, authenticators []authn.Authenticator, authorizers []authz.Authorizer) {
	authConfig := &AuthConfig{Config: config, Authenticators: authenticators, Authorizers: authorizers}
	session, err := mgo.DialWithInfo(&authConfig.Config.ACLMongoConf.DialInfo.DialInfo)
	if err != nil {
		//
	}

	userSession, err := mgo.DialWithInfo(&authConfig.Config.MongoAuth.DialInfo.DialInfo)
	if err != nil {
		//
	}
	ACManager = &AuthConfigManager{authConfig: authConfig, session: session, userSession: userSession}
}

// ======================================= 登陆相关=========================================
func (acm *AuthConfigManager) DoLogin(user string, password string) (string, bool) {
	for _, a := range acm.authConfig.Authenticators {
		result, err := a.Authenticate(user, authn.PasswordString(password))
		// glog.V(2).Infof("Authn %s %s -> %t, %s", a.Name(), ar.ai.Account, result, err)
		if err != nil {
			if err == authn.NoMatch {
				continue
			}
			// err = fmt.Errorf("authn #%d returned error: %s", i+1, err)
			// glog.Errorf("%s: %s", ar, err)
			return err.Error(), false
		}
		return "", result
	}
	// Deny by default.
	// glog.Warningf("%s did not match any authn rule", ar.ai)
	return "", false
}

/*
func (acm *AuthConfigManager) DoLogout(user string, password string) (string, bool) {
}
*/

// ======================================== 用户相关=========================================
func (acm *AuthConfigManager) AddUser(user string, password string) (string, bool) {
	// 这里password是已经加密的

	// 校验密码
	if !strings.HasPrefix(password, "$2y$05$") {
		return "密码格式不正确", false
	}

	collection := acm.userSession.DB(acm.authConfig.Config.MongoAuth.DialInfo.DialInfo.Database).C(acm.authConfig.Config.MongoAuth.Collection)
	err := collection.Insert(bson.M{"username": user, "password": password})
	if err != nil {
		return err.Error(), false
	}

	return "插入正常", true
}

func (acm *AuthConfigManager) DeleteUser(user string) (string, bool) {
	collection := acm.userSession.DB(acm.authConfig.Config.MongoAuth.DialInfo.DialInfo.Database).C(acm.authConfig.Config.MongoAuth.Collection)
	err := collection.Remove(bson.M{"Username": user})
	if err != nil {
		return err.Error(), false
	}

	return "删除正常", true
}

func (acm *AuthConfigManager) QueryAllUser() [][]string {
	// 查询出静态用户
	userMap := acm.authConfig.Config.Users
	users := make([][]string, len(userMap))
	i := 0
	for key, _ := range userMap {
		users[i] = []string{key, "0"}
		i = i + 1
	}

	sort.Sort(ByString(users))

	// 查询除mongo用户
	collection := acm.userSession.DB(acm.authConfig.Config.MongoAuth.DialInfo.DialInfo.Database).C(acm.authConfig.Config.MongoAuth.Collection)
	userInfos := []authn.UserInfo{}
	collection.Find(bson.M{}).All(&userInfos)

	for _, info := range userInfos {
		glog.Infoln(info.Username)
		glog.Infoln(info.Password)
		users = append(users, []string{info.Username, "1"})
	}
	return users
}

// ======================================= 镜像相关===========================================
var (
	ADD = 1
	DEL = 2
)

// 修改权限
func (acm *AuthConfigManager) Update(user string, name *string, mtype int, ispull bool) bool {
	// 细粒度的更新对应acl数据:
	// 增加 pull push
	// 删除 pull push
	acm.lock.Lock()
	defer acm.lock.Unlock() // 强行释放锁
	var ret bool
	if mtype == ADD {
		ret = acm.add(user, name, ispull)
	} else if mtype == DEL {
		ret = acm.del(user, name, ispull)
	} else {
		// do thing
	}

	acm.UpdateCache()
	return ret
}

func (acm *AuthConfigManager) add(user string, name *string, ispull bool) bool {
	collection := acm.session.DB(acm.authConfig.Config.ACLMongoConf.DialInfo.DialInfo.Database).C(acm.authConfig.Config.ACLMongoConf.Collection)
	var newACL authz.ACL
	var findRule bson.M

	if name == nil || *name == "" {
		findRule = bson.M{"match": bson.M{"account": user}}
	} else {
		findRule = bson.M{"match": bson.M{"account": user, "name": *name}}
	}
	collection.Find(findRule).All(&newACL)
	glog.Infof("match acl is %s", newACL)
	if len(newACL) == 0 {
		return false
	}

	matchAcl := newACL[0]

	var newadd string
	if ispull {
		newadd = "pull"
	} else {
		newadd = "push"
	}
	//
	if len(*(matchAcl.Actions)) == 0 {
		// 直接插入
		collection.Update(findRule, bson.M{"$set": bson.M{"actions": []string{newadd}}})
		return true
	}

	if len(*(matchAcl.Actions)) == 1 {
		if (*matchAcl.Actions)[0] != newadd {
			collection.Update(findRule, bson.M{"$set": bson.M{"actions": []string{"pull", "push"}}})
		}
	}

	return true
}

func (acm *AuthConfigManager) del(user string, name *string, ispull bool) bool {
	collection := acm.session.DB(acm.authConfig.Config.ACLMongoConf.DialInfo.DialInfo.Database).C(acm.authConfig.Config.ACLMongoConf.Collection)
	var newACL authz.ACL
	var findRule bson.M

	if name == nil || *name == "" {
		findRule = bson.M{"match": bson.M{"account": user}}
	} else {
		findRule = bson.M{"match": bson.M{"account": user, "name": *name}}
	}
	collection.Find(findRule).All(&newACL)

	glog.Infof("match acl is %s", newACL)
	if len(newACL) == 0 {
		return false
	}

	matchAcl := newACL[0]

	var todel string
	var toremain string
	if ispull {
		todel = "pull"
		toremain = "push"
	} else {
		todel = "push"
		toremain = "pull"
	}
	//
	if len(*(matchAcl.Actions)) == 0 {
		return false
	}

	if len(*(matchAcl.Actions)) == 1 {
		if (*matchAcl.Actions)[0] == todel {
			collection.Update(findRule, bson.M{"$set": bson.M{"actions": []string{}}})
			return true
		}
		return false
	}

	collection.Update(findRule, bson.M{"$set": bson.M{"actions": []string{toremain}}})

	return true
}

// 删除权限
func (acm *AuthConfigManager) Delete(user string, name string) (string, bool) {
	collection := acm.session.DB(acm.authConfig.Config.ACLMongoConf.DialInfo.DialInfo.Database).C(acm.authConfig.Config.ACLMongoConf.Collection)
	err := collection.Remove(bson.M{"match": bson.M{"account": user, "name": name}})
	if err != nil {
		return err.Error(), false
	}

	acm.UpdateCache()
	return "删除成功", true
}

// 添加权限
func (acm *AuthConfigManager) Add(user string, name string, actions []string) (string, bool) {
	collection := acm.session.DB(acm.authConfig.Config.ACLMongoConf.DialInfo.DialInfo.Database).C(acm.authConfig.Config.ACLMongoConf.Collection)
	err := collection.Insert(bson.M{"match": bson.M{"account": user, "name": name}, "actions": actions})
	if err != nil {
		return err.Error(), false
	}
	acm.UpdateCache()
	return "添加成功", true
}

// 更新mongodb中的缓存
func (acm *AuthConfigManager) UpdateCache() {
	// TODO 这里直接调用了authz包中的全局channel, 而在authz中读取channel进行异步重新
	// 加载, 有点破坏了面向对象的思想;
	authz.CH <- 1
}

// 查询所有权限
func (acm *AuthConfigManager) QueryAllAuth() [][]string {
	detail := make([][]string, 0)
	// 遍历所有权限,将用户名和user一致的权限返回
	for _, authorizer := range acm.authConfig.Authorizers {
		var canModify int
		acls := authorizer.GetAllACLs()
		if authorizer.Name() == "static ACL" {
			canModify = 0
		} else if authorizer.Name() == "MongoDB ACL" {
			canModify = 1
		}

		for _, acl := range acls {
			imgName := acl.Match.Name
			if imgName == nil {
				star := "*"
				imgName = &star
			}

			userName := acl.Match.Account
			actions := acl.Actions
			var ac = 0
			for _, action := range *actions {
				if action == "pull" {
					ac |= 1
				} else if action == "push" {
					ac |= 2
				} else if action == "*" {
					ac = 3
				}
			}
			detail = append(detail, []string{*userName, *imgName, strconv.Itoa(ac), strconv.Itoa(canModify)})
		}
	}

	fmt.Println(detail)
	return detail

}

// 查询某个用户所有的权限
func (acm *AuthConfigManager) QueryDetail(user string) [][]string {
	// detail 格式如下所示
	/* [
			["acl1", "xx", "yy"],
			["acl2", "xx", "yy"],
	   ]
	   // xx代表对于镜像的权限，xx 0代表没有权限，1代表只有拉取权限， 2代表只有推送权限
	   // 3代表拉取推送权限都有
	   // yy代表是否能够修改权限，对于static acl不能修改，只能通过修改文件重启auth才能修改
	   // 对于mongo acl可以直接操作。
	   // yy 0代表不可修改， 1代表可以修改
	*/
	detail := make([][]string, 0)
	// 遍历所有权限,将用户名和user一致的权限返回
	for _, authorizer := range acm.authConfig.Authorizers {
		var canModify int
		acls := authorizer.GetMatchACLs(user)
		if authorizer.Name() == "static ACL" {
			canModify = 0
		} else if authorizer.Name() == "MongoDB ACL" {
			canModify = 1
		}

		for _, acl := range acls {
			imgName := acl.Match.Name
			if imgName == nil {
				star := "*"
				imgName = &star
			}

			actions := acl.Actions
			var ac = 0
			for _, action := range *actions {
				if action == "pull" {
					ac |= 1
				} else if action == "push" {
					ac |= 2
				} else if action == "*" {
					ac = 3
				}
			}
			detail = append(detail, []string{*imgName, strconv.Itoa(ac), strconv.Itoa(canModify)})
		}
	}

	fmt.Println(detail)
	return detail
}
