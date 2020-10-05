package main

import (
	"log"
	"net/http"
	"os"

	//"test1/mvc/controller"
	"test1/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "2020"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//http.Handle("/", controller.Handler())
	srv := handler.NewDefaultServer(generated.NewExecutableSchema())
	http.Handle("/", srv)
	http.Handle("/play", playground.Handler("test", "/graphql"))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
