package main

import (
	"fmt"
	"github.com/gorilla/mux/internal/app/trelloserver"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func main() {
	cfg := trelloserver.NewConfig()

	f, err := os.Open("configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	fmt.Println(cfg)

}
