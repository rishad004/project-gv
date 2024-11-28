package repository

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rishad004/project-gv/user-service/internal/domain"
)

func (r *userRepo) PushCommentToQueue(topic string, message domain.Kafka) {
	msgByte, err := json.Marshal(message)
	if err != nil {
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msgByte),
	}

	partition, offset, err := r.Kfk.SendMessage(msg)
	if err != nil {
		return
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
