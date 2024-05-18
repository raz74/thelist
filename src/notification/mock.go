package notification

import (
	"context"
	"fmt"
	"log"
	"time"
)

type MockRepo struct {
	motifMap map[uint]*Notification
}

func (m MockRepo) Set(ctx context.Context, notification *Notification) error {
	log.Println("set", notification)
	m.motifMap[notification.Id] = notification
	time.Sleep(500 * time.Millisecond)
	return nil
}

func (m MockRepo) Get(ctx context.Context, id uint) (*Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockRepo) GetAll(ctx context.Context) ([]*Notification, error) {
	var results []*Notification
	for _, notification := range m.motifMap {
		results = append(results, notification)
	}
	return results, nil
}

func NewMockRepository() MockRepo {
	return MockRepo{motifMap: make(map[uint]*Notification)}
}

func (n Notification) GetMessage() string {
	return fmt.Sprintf("%+v", n.Message)
}

func (n Notification) GetId() uint {
	return n.Id
}
