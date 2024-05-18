package main

import (
	"TheList/config"
	"TheList/src/notification"
	"TheList/src/providers/databases"
	"TheList/src/providers/restapi"
	"TheList/src/restaurant"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("theList started ...")
	Init()
}

func Init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	logger := log.New(os.Stdout, "", 0)

	//init databases
	db, err := databases.ConnectPostgres()
	if err != nil {
		panic(err)
	}

	//init repositories
	restaurantRepo := restaurant.NewRepository(db)
	notificationRepo := notification.NewRepository(db)

	//init service
	notificationService := notification.NewService(notificationRepo)
	restaurantService := restaurant.NewService(restaurantRepo, notificationService, logger)

	ginEngine := restapi.InitGin(config.GetServerAddress())

	restaurant.InitHandler(ginEngine.GetEngine(), restaurantService)

	notificationService.Start()
	ginEngine.Start()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down service ...")

	ginEngine.Shutdown()
	notificationService.Shutdown()
}
