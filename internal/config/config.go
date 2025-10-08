package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type AppConfig struct {
	Name string
	Port string
	Env  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

type JWTConfig struct {
	Secret       string
	ExpiredHours int
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf(".env file not found, using environment variables")
	}

	config := &Config{
		App: AppConfig{
			Name: viper.GetString("APP_NAME"),
			Port: viper.GetString("APP_PORT"),
			Env:  viper.GetString("APP_ENV"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
			SSLMode:  viper.GetString("DB_SSLMODE"),
			TimeZone: viper.GetString("DB_TIMEZONE"),
		},
		JWT: JWTConfig{
			Secret:       viper.GetString("JWT_SECRET"),
			ExpiredHours: viper.GetInt("JWT_EXPIRED_HOURS"),
		},
	}

	if config.JWT.Secret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	if config.Database.Host == "" {
		return nil, fmt.Errorf("DB_HOST is required")
	}

	return config, nil
}
