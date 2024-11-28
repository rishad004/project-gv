package di

import (
	"sync"

	"github.com/rishad004/project-gv/notification-service/internal/handler"
)

func InitNotification() {
	var wg sync.WaitGroup

	connSub := InitSubscription()

	c := handler.NewConnections(connSub)

	wg.Add(3)
	go c.SubKafka(&wg)
	go c.ChatKafka(&wg)
	go c.WalletKafka(&wg)

	wg.Wait()
}
