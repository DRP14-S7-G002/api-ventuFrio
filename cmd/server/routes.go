package routes

import (
	"api-ventuFrio/cmd/server/config"
	"log"
	"net/http"
)

func HandlerRequests() {
	http.HandleFunc("/", config.Home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
