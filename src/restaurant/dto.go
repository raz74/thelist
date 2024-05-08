package restaurant

type CreateRestaurantReq struct {
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	FoodsType int     `json:"foods_type"`
}

type GetRestaurantReq struct {
	Id uint `param:"id"`
}

type RestaurantResp struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	FoodsType int     `json:"foods_type"`
}

func ToRestaurantDto(record *Restaurant) RestaurantResp {
	return RestaurantResp{
		Id:        record.ID,
		Name:      record.Name,
		Price:     record.Price,
		FoodsType: record.FoodsType,
	}
}
