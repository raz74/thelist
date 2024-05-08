package restaurant

import (
	"fmt"
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name      string
	Price     float32
	FoodsType int
}

func (r Restaurant) GetMessage() string {
	return fmt.Sprintf("%+v", r)
}

func (r Restaurant) GetId() uint {
	return r.ID
}
