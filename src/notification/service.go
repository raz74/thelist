package notification

import (
	ginErr "TheList/utilize/error"
	"context"
	"time"
)

const notifChanBufSize = 50

type ShutdownFunc func()

type Service struct {
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
	s.notifChan <- Notification{
		Id:      m.GetId(),
		Message: m.GetMessage(),
		Date:    time.Now(),
	}
}

func (s *Service) Start() ShutdownFunc {
	ctx := context.Background()
	go s.handle(ctx)
	return func() {
		ctx.Done()
	}
}

func (s *Service) handle(ctx context.Context) {
	for notifMessage := range s.notifChan {
		if err := s.handleMessage(ctx, notifMessage); err != nil {
			// todo
		}
	}
}

func (s *Service) handleMessage(ctx context.Context, message Notification) error {
	err := s.repo.Set(ctx, &message)
	if err != nil {
		return ginErr.SendNotificationFailedError // todo
	}
	return nil
}
