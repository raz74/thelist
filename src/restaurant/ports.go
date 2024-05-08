package restaurant

import "context"

type IRepository interface {
	Set(ctx context.Context, data *Restaurant) error
	Get(ctx context.Context, id uint) (*Restaurant, error)
}
