package restapi

import (
	ginErr "TheList/utilize/error"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type GinEngine struct {
	engine *gin.Engine
	srv    *http.Server
}

func InitGin(addr string) *GinEngine {
	engine := gin.Default()
	engine.Use(ginErr.ErrorHandler)

	srv := &http.Server{
		Addr:    addr,
		Handler: engine.Handler(),
	}

	return &GinEngine{
		engine: engine,
		srv:    srv,
	}
}

func (g *GinEngine) GetEngine() *gin.Engine {
	return g.engine
}

func (g *GinEngine) Start() {
	go func() {
		if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (g *GinEngine) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := g.srv.Shutdown(ctx); err != nil {
		log.Fatal("Gin Shutdown:", err)
	}
	log.Println("Gin exiting")
}
