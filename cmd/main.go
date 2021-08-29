package main

import (
	"log"

	"github.com/fbettic/ecommerce-back/internal/logs"
	"github.com/fbettic/ecommerce-back/rest"
)

func main() {
	_ = logs.InitLogger()
	log.Println("Starting API")
	rest.Start()
}
