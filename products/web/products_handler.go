package products

import (
	"../models"
	"encoding/json"
	"net/http"
)

var products = ProductsList{
	{
		ID:          1,
		Title:       "Producto 1",
		Description: "Descripcion 1",
		Price:       189.02,
	},
	{
		ID:          2,
		Title:       "Producto 2",
		Description: "Descripcion 2",
		Price:       2228.5,
	},
	{
		ID:          3,
		Title:       "Producto 3",
		Description: "Descripcion 3",
		Price:       1825.3,
	},
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products.Products)
}
