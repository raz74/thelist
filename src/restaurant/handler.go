package restaurant

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	service *Service
}

func InitHandler(gin *gin.Engine, service *Service) {
	handler := &Handler{service: service}
	handler.register(gin)
}

func (h Handler) register(gin *gin.Engine) {
	group := gin.Group("/restaurant")
	group.POST("/", h.create)
	group.GET("/:id", h.get)
}

func (h Handler) create(c *gin.Context) {
	var req CreateRestaurantReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, resp)
}

func (h Handler) get(c *gin.Context) {
	strId := c.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Get(c.Request.Context(), uint(intId))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, resp)
}
