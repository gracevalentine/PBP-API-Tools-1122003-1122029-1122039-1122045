package main

import (
	c "EksplorasiToolsAPI/Controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/reservation", c.HandleReservation).Methods("POST")
	router.HandleFunc("/reservation/show", c.ShowReservation).Methods("GET")

	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")

	log.Fatal(http.ListenAndServe(":8888", router))
}
