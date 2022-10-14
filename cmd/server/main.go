package main

import (
	"log"

	"github.com/franciscoalface/golang-product-api/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	log.Println(config)
	if err != nil {
		panic(err)
	}
}
