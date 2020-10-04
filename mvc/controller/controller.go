package controller

import (
	"net/http"
	"strings"

	//"test1/mvc/controller/client1"

	//"test1/mvc/model"
	"test1/mvc/view"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimLeft(r.RequestURI, "/") {
		case "client1":
			//client1.Handler(w, r)
			view.Client1(w, r)
		default:
			view.PageNotFound(w, r)
		}
	}
}
