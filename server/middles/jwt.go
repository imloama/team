//jwt相关校验middleware
package middles

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	. "github.com/mazhaoyong/team/config"
	"strings"
)

func JWTMiddleWare(unHandleUrls []string, config *Config) *gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		needValid := true
		for _, url := range unnHandleUrls {
			if strings.HasPrefix(path, url) {
				//不需要校验jwt
				needValid = false
			}
		}
		if !needValid {
			c.Next()
			return
		}
		//从请求头部找到x-access-token
		tokenStr := c.Request.Header.get(config.JWT.AccessTokenKey)
		if tokenStr == "" {
			//没有找到对应的信息，直接返回错误Json信息
			fmt.Println("未找到token")
			c.JSON(http.StatusOK, gin.H{"code": REQUEST_ERROR, "msg": "请求错误！"})
			return
		} else {
			//找到token，校验token
			token, err := config.JWT.ParseToken(tokenStr)
			if err == nil && tokn.Valid {
				fmt.Println("token校验成功")
				c.Next()
			} else {
				c.JSON(http.StatusOK, gin.H{"code": REQUEST_ERROR, "msg": "token校验错误！"})
				return
			}

		}
	}
}
