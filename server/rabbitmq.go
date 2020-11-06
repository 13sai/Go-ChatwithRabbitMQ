package server

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	
	"local.com/sai0556/Go-ChatwithRabbitMQ/util"
)

type RabbitMQ struct {
	Client *amqp.Connection
}

func (server *RabbitMQ) getClient() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", viper.GetString("rabbit.user"), viper.GetString("rabbit.pw"), viper.GetString("rabbit.host"), viper.GetString("rabbit.port")))
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	server.Client = conn
}

