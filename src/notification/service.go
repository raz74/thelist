package notification

import (
	ginErr "TheList/utilize/error"
	"context"
	"log"
	"sync"
	"time"
)

const (
	notifChanBufSize = 50
	numHandlers      = 5
)

type ShutdownFunc func()

type Service struct {
	sync.WaitGroup
	notifChan chan Notification
	repo      IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo, notifChan: make(chan Notification, notifChanBufSize)}
}

func (s *Service) AsyncSend(m IMessage) {
	go s.Send(m)
}

func (s *Service) Send(m IMessage) {
	defer func() { recover() }()
	s.notifChan <- Notification{
		Id:      m.GetId(),
		Message: m.GetMessage(),
		Date:    time.Now(),
	}
}

func (s *Service) Start() {
	ctx := context.Background()
	for i := 0; i < numHandlers; i++ {
		go s.handle(ctx)
		s.Add(1)
	}
}

func (s *Service) Shutdown() {
	log.Println("exiting notification service")
	close(s.notifChan)
	s.Wait()
	log.Println("notification service exited")
}

func (s *Service) handle(ctx context.Context) {
	for notifMessage := range s.notifChan {
		if err := s.handleMessage(ctx, notifMessage); err != nil {
			log.Println("err in handle notification message:", err)
		}
	}
	s.Done()
}

func (s *Service) handleMessage(ctx context.Context, message Notification) error {
	err := s.repo.Set(ctx, &message)
	if err != nil {
		return ginErr.SendNotificationFailedError // todo
	}
	return nil
}
