package main

import (
	"ecommerce-back/internal/logs"
	"ecommerce-back/rest"
	"log"
)

func main() {
	_ = logs.InitLogger()
	log.Println("Starting API")
	rest.Start(":3001")
}
