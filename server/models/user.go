// 用户
package models

import (
	. "github.com/mazhaoyong/team/server/config"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Code     string        //登陆名
	Name     string        //显示名
	Password string        //密码
	Email    string        //邮箱
}

const USER_COLLECTION = "sys_user"

func (this *User) Add(config *Config) error {
	session := config.Mongo.GetSession()
	defer session.Close()
	c := session.DB(config.Mongo.Name).C(USER_COLLECTION)
	return c.Insert(entity)
}
