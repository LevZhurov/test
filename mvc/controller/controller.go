package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"test1/mvc/model"
	"test1/mvc/view"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimLeft(r.RequestURI, "/") {
		case "client1":
			view.Client1(w, r)
		case "graphql":
			json.NewEncoder(w).Encode(model.GraphQL(r.URL.Query().Get("query")))
		default:
			view.PageNotFound(w, r)
		}
	}
}
