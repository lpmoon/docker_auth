package authn

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// the user and encrypted password are stored in the MongoDB
// admin has full access to all data.
// 1) add new user
// 2) delete old user
// 3) query user and authenticate
// 4) query all users
type MongoAuthConfig struct {
	DialInfo   MongoAuthDialConfig `yaml:"dial_info,omitempty"`
	Collection string              `yaml:"collection,omitempty"`
	CacheTTL   time.Duration       `yaml:"cache_ttl,omitempty"`
}

type MongoAuthDialConfig struct {
	mgo.DialInfo `yaml:",inline"`
	PasswordFile string `yaml:"password_file,omitempty"`
}

type MongoAuth struct {
	config  *MongoAuthConfig
	session *mgo.Session
}

type UserInfo struct {
	Username string         `yaml:"Username,omitempty" json:"Username,omitempty"`
	Password PasswordString `yaml:"Password,omitempty" json:"Password,omitempty"`
}

func NewMongoAuth(c *MongoAuthConfig) (*MongoAuth, error) {
	// get session
	session, err := mgo.DialWithInfo(&c.DialInfo.DialInfo)
	if err != nil {
		return nil, err
	}
	mongoAuth := &MongoAuth{
		config:  c,
		session: session,
	}

	return mongoAuth, nil

}

func (ma *MongoAuthConfig) Validate() error {

	if len(ma.DialInfo.DialInfo.Addrs) == 0 {
		return errors.New("At least one element in acl_mongo.dial_info.addrs is required")
	}
	if ma.DialInfo.DialInfo.Timeout == 0 {
		ma.DialInfo.DialInfo.Timeout = 10 * time.Second
	}
	if ma.DialInfo.DialInfo.Database == "" {
		return errors.New("acl_mongo.dial_info.database is required")
	}
	if ma.Collection == "" {
		return errors.New("acl_mongo.collection is required")
	}
	if ma.CacheTTL < 0 {
		return errors.New(`acl_mongo.cache_ttl is required (e.g. "1m" for 1 minute)`)
	}
	return nil

}

func (ma *MongoAuth) Authenticate(user string, password PasswordString) (bool, error) {
	// 获取连接
	collection := ma.session.DB(ma.config.DialInfo.DialInfo.Database).C(ma.config.Collection)
	var userInfos []UserInfo
	collection.Find(bson.M{"user": user}).All(&userInfos)

	if len(userInfos) == 0 {
		return false, nil
	}

	// 取第一个
	tocompare := userInfos[0]
	if bcrypt.CompareHashAndPassword([]byte(tocompare.Password), []byte(password)) != nil {
		return false, nil
	}

	return true, nil

}

func (ma *MongoAuth) Stop() {

}

func (ma *MongoAuth) Name() string {
	return "MONGO"
}
