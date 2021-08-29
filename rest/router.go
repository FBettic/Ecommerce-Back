package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/fbettic/ecommerce-back/internal/logs"
	"github.com/fbettic/ecommerce-back/products"

	"github.com/gorilla/mux"
)

const AllowedCORSDomain = "http://localhost"

func Router(port string) {
	router := mux.NewRouter().StrictSlash(true)

	enableCORS(router)

	router.HandleFunc("/products", getProductHandler).Methods("GET")
	router.HandleFunc("/products", createProductHandler).Methods("POST")
	router.HandleFunc("/products/{id}", getProductHandlerById).Methods("GET")
	router.HandleFunc("/products/{id}", updateProductHandler).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProductHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	products, err := productsHandler.GetProducts()

	if err == nil {
		respondWithSuccess(products, w)
	} else {
		respondWithError(err, w)
	}
}

func getProductHandlerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := stringToInt64(vars["id"])

	if err != nil {
		logs.Log().Errorf("invalid ID: %s", err.Error())
		respondWithError(err, w)
		return
	}

	product, err := productsHandler.GetProductById(id)

	if err != nil {
		respondWithError(err, w)
	} else {
		respondWithSuccess(product, w)
	}
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Log().Errorf("cannot read the body: %s", err.Error())
		return
	}

	product := products.Product{}
	json.Unmarshal(reqBody, &product)

	if err != nil {
		respondWithError(err, w)
	} else {
		err := productsHandler.CreateProduct(&product)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := stringToInt64(vars["id"])
	if err != nil {
		logs.Log().Errorf("invalid ID: %s", err.Error())
		respondWithError(err, w)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Log().Errorf("cannot read the body: %s", err.Error())
		respondWithError(err, w)
		return
	}

	product := products.Product{}
	json.Unmarshal(reqBody, &product)

	if err != nil {
		respondWithError(err, w)
	} else {
		err := productsHandler.UpdateProduct(&product, id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := stringToInt64(vars["id"])
	if err != nil {
		logs.Log().Errorf("Invalid ID: %s", err.Error())
		respondWithError(err, w)
		return
	}

	err = productsHandler.DeleteProductById(id)

	if err != nil {
		respondWithError(err, w)
	} else {
		respondWithSuccess(true, w)
	}
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func stringToInt64(s string) (int64, error) {
	num, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return num, err
}
