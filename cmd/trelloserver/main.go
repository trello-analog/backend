package main

import (
	"github.com/trello-analog/backend/server"
	"log"
)

func main() {
	app := server.NewApp()

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
