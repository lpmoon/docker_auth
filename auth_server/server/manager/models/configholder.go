package models

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
	"sync"
)

type ByString []string

func (s ByString) Len() int {
	return len(s)
}

func (s ByString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByString) Less(i, j int) bool {
	return s[i] < s[j]
}

type AuthConfig struct {
	Config         *config.Config
	Authenticators []authn.Authenticator
	Authorizers    []authz.Authorizer
}

type AuthConfigManager struct {
	authConfig *AuthConfig
	session    *mgo.Session
	lock       sync.Mutex
}

var ACManager *AuthConfigManager

func InitAuthConfigManager(config *config.Config, authenticators []authn.Authenticator, authorizers []authz.Authorizer) {
	authConfig := &AuthConfig{Config: config, Authenticators: authenticators, Authorizers: authorizers}
	session, err := mgo.DialWithInfo(&authConfig.Config.ACLMongoConf.DialInfo.DialInfo)
	if err != nil {
		//
	}
	ACManager = &AuthConfigManager{authConfig: authConfig, session: session}
}

func (acm *AuthConfigManager) QueryAllUser() []string {
	userMap := acm.authConfig.Config.Users
	users := make([]string, len(userMap))
	i := 0
	for key, _ := range userMap {
		users[i] = key
		i = i + 1
	}

	sort.Sort(ByString(users))

	return users
}

var (
	ADD = 1
	DEL = 2
)

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

func (acm *AuthConfigManager) UpdateCache() {
	// 更新
	authz.CH <- 1
}

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
