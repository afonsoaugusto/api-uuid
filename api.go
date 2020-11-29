package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
)

// func helloHandler(w http.ResponseWriter, req *http.Request) {
//         fmt.Fprintf(w, "Hello, %s!\n", mux.Vars(req)["name"])
// }

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
}

func generateUuid(w http.ResponseWriter, req *http.Request) {
	id := uuid.Must(uuid.NewRandom())
	fmt.Fprintf(w, id.String())
}

func main() {
	router := mux.NewRouter()
	router.Use(apmgorilla.Middleware())

	// router.HandleFunc("/hello/{name}", helloHandler).Methods("GET")
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/", generateUuid).Methods("GET")

	// port
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
