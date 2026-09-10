package main

import (
	"multispore/internal/server"
	"multispore/internal/utils"
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	envFile, err := godotenv.Read(".env")

	hasConfig := err == nil

	if !hasConfig {
		log.Warnf("Cannot find .env file, using default!")
	}

	srv, err := server.New(
		// This is the server config
		&server.ServerConfig{
			Host: utils.If(hasConfig, envFile["SERVER_HOST"], "127.0.0.1"),
			Port: utils.If(hasConfig, envFile["SERVER_PORT"], "5523"),
		},
	)

	if err != nil {
		log.Errorf("Failed to create the server object: %v", err)
		os.Exit(1)
	} else {
		srv.Run()
	}
}
