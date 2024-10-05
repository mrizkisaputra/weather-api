package main

import (
	"github.com/joho/godotenv"
	"wheater-api/config"
)

func main() {
	app := config.NewApp()
	log := config.NewLogger()

	if err := godotenv.Load("./config.env"); err != nil {
		log.WithError(err).Panic("Error loading .env file")
	}

	config.NewBootstrap(&config.Bootstrap{
		App: app,
		Log: log,
	})

	panic(app.Listen(":8080"))
}
