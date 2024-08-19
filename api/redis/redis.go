package redis

import (
	"context"
	"gateway/internal/config"

	// pb "reservation_service/genproto/reservation"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func connectDB(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: "",
		DB:       0,
	})

	return rdb
}

func SaveVerificationCode(config *config.Config, email string, code string, expiration time.Duration) error {
	client := connectDB(config)

	ctx := context.Background()
	err := client.Set(ctx, email, code, expiration).Err()
	if err != nil {
		return errors.Wrap(err, "failed to save verification code in redis")
	}
	return nil
}

func GetVerificationCode(config *config.Config, email string) (string, error) {
	client := connectDB(config)

	ctx := context.Background()
	code, err := client.Get(ctx, email).Result()
	if err == redis.Nil {
		return "", errors.New("no verification code found for email")
	} else if err != nil {
		return "", errors.Wrap(err, "failed to get verification code from redis")
	}
	return code, nil
}
