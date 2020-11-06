package main

import (
	"github.com/spf13/pflag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"local.com/sai0556/Go-ChatwithRabbitMQ/config"
	"local.com/sai0556/Go-ChatwithRabbitMQ/router"
	"local.com/sai0556/Go-ChatwithRabbitMQ/server"
)

var (
	conf = pflag.StringP("config", "c", "", "config filepath")
)

func main() {
	pflag.Parse()

	// 初始化配置
	if err := config.Init(*conf); err != nil {
		panic(err)
	}

	ws := server.WebSocket{}
	go ws.Start()

	gin.SetMode(viper.GetString("runMode"))
	g := gin.New()
	g = router.Load(g)
	g.Run(viper.GetString("addr"))
}