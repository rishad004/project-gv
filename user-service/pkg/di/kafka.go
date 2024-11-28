package di

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

func InitKafka() sarama.SyncProducer {
	brokersUrl := []string{viper.GetString("KAFKA_PORT")}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		log.Fatal("unable to connect kafka producer : ", err)
	}

	return conn
}
