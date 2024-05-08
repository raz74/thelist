package notification

import (
	"TheList/src/providers/databases"
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	databases.MustMigrate(db, Notification{})
	return &Repository{db: db}
}

func (r *Repository) Set(ctx context.Context, notification *Notification) error {
	return r.db.WithContext(ctx).Create(&notification).Error
}

func (r *Repository) Get(ctx context.Context, id uint) (*Notification, error) {
	var notif Notification
	err := r.db.WithContext(ctx).Where("id=?", id).First(&notif).Error
	return &notif, err
}

func (r *Repository) GetAll(ctx context.Context) ([]*Notification, error) {
	var notif []*Notification
	err := r.db.WithContext(ctx).Find(&notif).Error
	return notif, err
}
