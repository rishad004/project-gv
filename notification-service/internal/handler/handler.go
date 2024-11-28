package handler

import "github.com/IBM/sarama"

type Connections struct {
	SubConn sarama.Consumer
}

func NewConnections(s sarama.Consumer) *Connections {
	return &Connections{SubConn: s}
}
