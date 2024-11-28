package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/rishad004/project-gv/notification-service/internal/domain"
	"github.com/rishad004/project-gv/notification-service/utils"
)

func (h *Connections) SubKafka(wg *sync.WaitGroup) {
	var kafka domain.Kafka

	defer wg.Done()

	consumer, err := h.SubConn.ConsumePartition("Subscription", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				if err = json.Unmarshal(msg.Value, &kafka); err != nil {
					log.Println(err)
				}

				data := strings.Split(string(kafka.Message), ",")

				if err = utils.SendEmail(data[0], "Subscription Successful!", "Dear User,\nYour subscription of "+data[1]+" has been successfully activated!\nThank you for supporting Streamer and being part of their amazing community.\n\nBest regards,\nGamer Vision"); err != nil {
					log.Println(err)
				}

				fmt.Printf("Received message: Topic(%s) | Message(%s) \n", string(msg.Topic), string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Subscription Consumer Closed")

	if err := h.SubConn.Close(); err != nil {
		panic(err)
	}
}

func (h *Connections) ChatKafka(wg *sync.WaitGroup) {
	var kafka domain.Kafka

	defer wg.Done()

	consumer, err := h.SubConn.ConsumePartition("Superchat", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				if err = json.Unmarshal(msg.Value, &kafka); err != nil {
					log.Println(err)
				}

				data := strings.Split(string(kafka.Message), ",")

				if err = utils.SendEmail(data[0], "Your Superchat Was Sent Successfully!", "Dear User,\nThank you for your Superchat during the live session!\nYour superchat of "+data[1]+" has been successfully sent and is now visible in the live chat for the community and host to see.\n\n!Best regards,\nGamer Vision"); err != nil {
					log.Println(err)
				}

				fmt.Printf("Received message: Topic(%s) | Message(%s) \n", string(msg.Topic), string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Superchat Consumer Closed")

	if err := h.SubConn.Close(); err != nil {
		panic(err)
	}
}

func (h *Connections) WalletKafka(wg *sync.WaitGroup) {
	var kafka domain.Kafka

	defer wg.Done()

	consumer, err := h.SubConn.ConsumePartition("WalletAdd", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				if err = json.Unmarshal(msg.Value, &kafka); err != nil {
					log.Println(err)
				}

				data := strings.Split(string(kafka.Message), ",")

				if err = utils.SendEmail(data[0], "Deposit: Wallet Updated", "Dear User,\nYour deposit of "+data[1]+" has been successfully processed and added to your wallet.\nThank you for choosing Gamer Vision\n\n!Best regards,\nGamer Vision"); err != nil {
					log.Println(err)
				}

				fmt.Printf("Received message: Topic(%s) | Message(%s) \n", string(msg.Topic), string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("WalletAdd Consumer Closed")

	if err := h.SubConn.Close(); err != nil {
		panic(err)
	}
}
