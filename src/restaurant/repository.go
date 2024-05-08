package restaurant

import (
	"TheList/src/providers/databases"
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	databases.MustMigrate(db, Restaurant{})
	return &Repository{db: db}
}

func (p Repository) Set(ctx context.Context, data *Restaurant) error {
	return p.db.WithContext(ctx).Create(&data).Error
}

func (p Repository) Get(ctx context.Context, id uint) (*Restaurant, error) {
	var data Restaurant
	err := p.db.WithContext(ctx).Where("id =?", id).First(&data).Error
	return &data, err
}

// todo : write get all and delete
