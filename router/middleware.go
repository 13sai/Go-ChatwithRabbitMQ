package router

import (
	"net/http"
	"strings"
	"fmt"

	"github.com/gin-gonic/gin"

	"local.com/sai0556/Go-ChatwithRabbitMQ/server"
)


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
 
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
 
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		// 处理请求
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		token := c.Query("token")
		if len(auth) == 0 && len(token) == 0 {
			server.SendResponse(c, http.StatusUnauthorized, "未登录", auth)
			return 
		}

		if len(token) > 0 {
			auth = token
		} else {
			auth = strings.Fields(auth)[1]
		}

		
		// 校验token
		_, Claims, err := server.ParseToken(auth)
		if err != nil {
			fmt.Println(err)
			server.SendResponse(c, http.StatusUnauthorized, "未登录", nil)
			return 
		}
		fmt.Println(Claims)
		c.Set("userId", Claims.Id)
		c.Set("avatar", Claims.Avatar)
		c.Set("nickname", Claims.Username)

		c.Next()
	}
}