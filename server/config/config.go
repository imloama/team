// 配置文件
package config

import (
	"flag"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rakyll/globalconf"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CApp struct {
	Name    string
	Host    string
	Port    int
	Version string
}

type CMongo struct {
	Url     string
	Name    string
	session *mgo.Session // mongodb数据库连接
}

type CJWT struct {
	SharedSecret   string
	AccessTokenKey string
}

type Config struct {
	App   *CApp
	Mongo *CMongo
	JWT   *CJWT
}

const CONFIG_FILE = "./app.conf"

const SHARED_SECRET = "!@#$%^&*()_+1234567890-="
const ACCESS_TOKEN_KEY = "x-access-token"

const (
	REQUEST_ERROR = 401
	UN_AUTH       = 403
	OK            = 200
	SERVER_ERROR  = 500
)

//初始化配置文件
func InitConfig() (*Config, error) {
	var appName = flag.String("app.name", "team", "应用名称")
	var appPort = flag.Int("app.port", 8080, "应用端口")
	var appHost = flag.String("app.host", "127.0.0.1", "应用主机")
	var appVersion = flag.String("app.version", "1.0.0", "应用版本")
	var mongoUrl = flag.String("mongo.url", "127.0.0.1:27017", "mongodb连接地址")
	var mongoName = flag.String("mongo.name", "x", "mongodb连接数据库")

	conf, err := globalconf.NewWithOptions(&globalconf.Options{
		Filename: CONFIG_FILE,
	})
	if err != nil {
		fmt.Printf("read config file error:%v\n", err)
		return nil, err
	}
	conf.ParseAll()
	app := &CApp{
		Name:    *appName,
		Host:    *appHost,
		Port:    *appPort,
		Version: *appVersion,
	}
	mongo := &CMongo{
		Url:  *mongoUrl,
		Name: *mongoName,
	}
	jwt := &CJWT{
		SharedSecret:    SHARED_SECRET,
		AccessTokenKeyh: ACCESS_TOKEN_KEY,
	}
	return &Config{
		App:   app,
		Mongo: mongo,
		JWT:   jwt,
	}, nil
}

func (this *CMongo) GetSession() *mgo.Session {
	if this.session == nil {
		var err error
		this.session, err = mgo.Dial(this.Url)
		this.session.SetMode(mgo.Monotonic, true)
		if err != nil {
			panic(err) //直接终止程序运行
		}
	}
	//最大连接池默认为4096
	return this.session.Clone()
}

/*
//公共方法，获取collection对象
func (this *CMongo) witchCollection(collection string, s func(*mgo.Collection) error) error {
	session := this.getSession()
	defer session.Close()
	c := session.DB(this.Name).C(collection)
	return s(c)
}

// 新增
func (this *CMongo) Insert(collection string, entity interface{}) error {
	query := func(c *mgo.Collection) error {
		return c.Insert(entity)
	}
	err := witchCollection(collection, query)
	if err != nil {
		return err
	}
	return nil
}

*/

func (this *CJWT) ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return this.SharedSecret, nil
	})
}

func (this *CJWT) NewToken(subject interface{}) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["subject"] = subject
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	return token.SignedString(this.SharedSecret)
}
