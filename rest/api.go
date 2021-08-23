package rest

import (
	"ecommerce-back/internal/database"
	"ecommerce-back/internal/logs"
)

func Start(port string) {

	client := database.NewSqlClient("root:root@/")
}
