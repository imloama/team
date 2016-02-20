// 首界面相关，登陆，主界面数据
package routers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	. "github.com/mazhaoyong/team/config"
)

func Index(c *gin.Context) {

}

func Login(c *gin.Context, config *Config) {
	config.JWT.NewToken("user")
}
