package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page!")
}

func HandlerRequests() {
	http.HandleFunc("/", Home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	HandlerRequests()

}
