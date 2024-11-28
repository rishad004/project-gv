package di

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

func InitSubscription() sarama.Consumer {
	brokersUrl := []string{viper.GetString("KAFKA_PORT")}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
