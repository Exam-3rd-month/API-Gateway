package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	ResetPassword ResetPasswordConfig
}

type ServerConfig struct {
	GATEWAY_PORT string
	Order_Port string
	Auth_Port  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type ResetPasswordConfig struct {
	From string
	Password string
}

func (c *Config) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	c.Server.GATEWAY_PORT = ":" + os.Getenv("GATEWAY_PORT")
	c.Server.Order_Port = ":" + os.Getenv("ORDER_PORT")
	c.Server.Auth_Port = ":" + os.Getenv("AUTH_PORT")
	c.Database.Host = os.Getenv("DB_HOST")
	c.Database.Port = os.Getenv("DB_PORT")
	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.DBName = os.Getenv("DB_NAME")

	c.ResetPassword.From = os.Getenv("RESET_PASSWORD_FROM")
	c.ResetPassword.Password = os.Getenv("EMAIL_PASSWORD")

	return nil
}

func New() (*Config, error) {
	config := &Config{}
	err := config.Load()
	if err != nil {
		return nil, err
	}
	return config, nil
}
