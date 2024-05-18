package notification

import (
	"context"
	"github.com/magiconair/properties/assert"
	"log"
	"testing"
	"time"
)

func Test_gracefulShutdown(t *testing.T) {
	//var notificationsChan = make(chan Notification, notifChanBufSize)
	//go func() {
	//	for i := 1; i < 5; i++ {
	//		notificationsChan <- Notification{
	//			Id:      uint(i),
	//			Message: "",
	//			Date:    time.Now(),
	//		}
	//		time.Sleep(500 * time.Millisecond)
	//	}
	//}()

	repo := NewMockRepository()
	service := NewService(repo)

	//shutdown := service.Start()

	for i := 1; i <= 5; i++ {
		service.Send(Notification{
			Id:      uint(i),
			Message: "",
			Date:    time.Now(),
		})
	}

	service.Start()

	//shutdown := service.Start()

	//go func() {
	//	for notif := range notificationsChan {
	//		service.Send(notif)
	//	}
	//}()
	//time.Sleep(3 * time.Second)
	service.Shutdown()

	result, err := service.repo.GetAll(context.Background())
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, len(result), 5)
}
