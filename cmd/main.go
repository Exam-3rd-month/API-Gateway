package main

import (
	"gateway/api"
	"gateway/internal/config"
	"log"
	"log/slog"
	"os"
)

func main() {
	logFile, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln("Error opening log file:", err)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	api := api.New()

	log.Fatal(api.NewRouter(configs, logger).Run("gateway" + configs.Server.GATEWAY_PORT))
}
