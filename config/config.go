package config

import "github.com/spf13/viper"

var cfg config

type config struct {
	ServerAddress string   `mapstructure:"SERVER_ADDRESS"`
	Postgres      postgres `mapstructure:",squash"`
}

type postgres struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     int    `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DbName   string `mapstructure:"POSTGRES_DB"`
}

func LoadConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&cfg)
	return nil
}

func GetServerAddress() string {
	return cfg.ServerAddress
}

func GetPostgres() *postgres {
	return &cfg.Postgres
}
