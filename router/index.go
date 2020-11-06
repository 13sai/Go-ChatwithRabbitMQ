package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"local.com/sai0556/Go-ChatwithRabbitMQ/server"
)

func Load(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())

	g.NoRoute(func (c *gin.Context)  {
		c.String(http.StatusNotFound, "404 not found");
	})

	// g.Use(gin.BasicAuth(gin.Accounts{
    //     "admin": "123456",
	// }))
	
	g.Use(Cors())
	g.POST("/login", server.Login)
	g.GET("/user/info", server.Info)

	return g
}