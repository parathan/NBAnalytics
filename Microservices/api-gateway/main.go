package main

import (
	"api-gateway/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/teams/twoteams", controller.TwoteamsController)
	log.Fatal(http.ListenAndServe(":8080", nil))
}