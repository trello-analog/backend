package main

import (
	"log"

	"github.com/trello-analog/backend/server"
)

func main() {
	app := server.NewApp()

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
