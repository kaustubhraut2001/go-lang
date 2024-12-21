package main // package main
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("mux in go lang")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func greeet() {
	fmt.Println("hello world")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serve home")
	w.Write([]byte("<h1> Server home <h1>"))
}
