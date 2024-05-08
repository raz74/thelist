package restaurant

import (
	"TheList/src/notification"
	ginErr "TheList/utilize/error"
	"context"
	"log"
)

type Service struct {
	repo                IRepository
	logger              *log.Logger
	notificationService notification.IService
}

func NewService(repo IRepository, notificationService notification.IService, logger *log.Logger) *Service {
	return &Service{repo: repo, logger: logger, notificationService: notificationService}
}

func (r Service) Create(ctx context.Context, req CreateRestaurantReq) (RestaurantResp, error) {
	var restaurant = &Restaurant{
		Name:      req.Name,
		Price:     req.Price,
		FoodsType: req.FoodsType,
	}
	err := r.repo.Set(ctx, restaurant)
	if err != nil {
		//todo handle errors
		r.logger.Printf("")
		return RestaurantResp{}, ginErr.RestaurantDontCreatedError
	}

	r.notificationService.AsyncSend(restaurant)

	return ToRestaurantDto(restaurant), nil
}

func (r Service) Get(ctx context.Context, id uint) (RestaurantResp, error) {
	restaurant, err := r.repo.Get(ctx, id)
	if err != nil {
		return RestaurantResp{}, ginErr.RestaurantNotFoundError
	}
	return ToRestaurantDto(restaurant), nil
}
