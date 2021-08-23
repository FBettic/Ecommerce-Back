package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", GetProducts).Methods("GET")
	/*
		router.HandleFunc("/tasks", getTasks).Methods("GET")
		router.HandleFunc("/tasks", createTask).Methods("POST")
		router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
		router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
		router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	*/

	log.Fatal(http.ListenAndServe(":3001", router))
}
