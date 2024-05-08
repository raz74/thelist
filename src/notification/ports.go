package notification

import "context"

type IRepository interface {
	Set(ctx context.Context, notification *Notification) error
	Get(ctx context.Context, id uint) (*Notification, error)
	GetAll(ctx context.Context) ([]*Notification, error)
}

type IMessage interface {
	GetMessage() string
	GetId() uint
}

type IService interface {
	AsyncSend(m IMessage)
	Send(m IMessage)
}
