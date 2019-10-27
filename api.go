package main                                                 
import (                                                     
        "fmt"                                                
        "log"                                                
        "net/http"                                           
		"github.com/gorilla/mux"
		"go.elastic.co/apm/module/apmgorilla"
		"github.com/google/uuid"
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
		
        log.Fatal(http.ListenAndServe(":8080", router))           
}