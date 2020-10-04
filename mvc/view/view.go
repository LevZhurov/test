package view

import (
	"html/template"
	"net/http"
)

func Client1(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("./mvc/view/html/client1_view.html"))
	tpl.Execute(w, nil)
}

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("./mvc/view/html/404_view.html"))
	tpl.Execute(w, nil)
}
