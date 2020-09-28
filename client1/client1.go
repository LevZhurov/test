package client1

import (
	"html/template"
	"net/http"
	//"os"
)

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "2021"
// 	}

// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/client1", client1Handler)
// 	http.ListenAndServe(":"+port, mux)
// }

var tpl = template.Must(template.ParseFiles("client1/index.html"))

func Client1Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Write(view())
		tpl.Execute(w, nil)
	}
}

func view() []byte {
	return []byte("")
}
