package main

import (
	"log"
	"net/http"
	"os"

	"test1/mvc/controller"
)

const defaultPort = "2020"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", controller.Handler())

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
