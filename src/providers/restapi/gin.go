package restapi

import (
	ginErr "TheList/utilize/error"
	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	engine *gin.Engine
	addr   string
}

func InitGin(addr string) *GinEngine {
	engine := gin.Default()
	engine.Use(ginErr.ErrorHandler)
	return &GinEngine{
		engine: engine,
		addr:   addr,
	}
}

func (g *GinEngine) GetEngine() *gin.Engine {
	return g.engine
}

func (g *GinEngine) Start() error {
	return g.engine.Run(g.addr)
}
