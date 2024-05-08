package databases

import (
	"TheList/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	dns := getPostgresDns()
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

func getPostgresDns() string {
	cfg := config.GetPostgres()
	dns := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)
	return dns
}

func MustMigrate(db *gorm.DB, entities ...interface{}) {
	if err := db.AutoMigrate(entities...); err != nil {
		panic(err)
	}
}
