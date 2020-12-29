package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloworld(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")

}

func handleRequest() {

	myroute := mux.NewRouter().StrictSlash(true)
	myroute.HandleFunc("/", helloworld).Methods("GET")
	myroute.HandleFunc("/practical", AllUser).Methods("GET")
	myroute.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myroute.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myroute.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", myroute))

}

func main() {

	fmt.Println("go orm tutoiral")

	InitialMigration()

	handleRequest()

}
