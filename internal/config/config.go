package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server        ServerConfig
	Redis         RedisConfig
	ResetPassword ResetPasswordConfig
}

type ServerConfig struct {
	GATEWAY_PORT string
	Order_Port   string
	Auth_Port    string
}

type ResetPasswordConfig struct {
	From     string
	Password string
}
type RedisConfig struct {
	Host string
	Port string
}

func (c *Config) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	c.Server.GATEWAY_PORT = ":" + os.Getenv("GATEWAY_PORT")
	c.Server.Order_Port = ":" + os.Getenv("ORDER_PORT")
	c.Server.Auth_Port = ":" + os.Getenv("AUTH_PORT")
	c.Redis.Host = os.Getenv("REDIS_HOST")
	c.Redis.Port = os.Getenv("REDIS_PORT")
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
