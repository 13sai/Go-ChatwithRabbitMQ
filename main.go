package main

import (
	"os"
	"log"
	"errors"

	"github.com/spf13/pflag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"local.com/sai0556/Go-ChatwithRabbitMQ/config"
	"local.com/sai0556/Go-ChatwithRabbitMQ/router"
	"local.com/sai0556/Go-ChatwithRabbitMQ/server"
	"local.com/sai0556/Go-ChatwithRabbitMQ/model"
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

	// 连接mysql数据库
	btn := model.GetInstance().InitPool()
	if !btn {
		log.Println("init db pool failure...")
		panic(errors.New("init db pool failure"))
		os.Exit(1)
	}

	// redis
	model.RedisInit()

	ws := server.WebSocket{}
	go ws.Start()

	gin.SetMode(viper.GetString("runMode"))
	g := gin.New()
	g = router.Load(g)
	g.Run(viper.GetString("addr"))
}