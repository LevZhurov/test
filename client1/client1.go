package client1

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("client1/index.html"))

func Client1Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func view() []byte {
	return []byte("")
}
