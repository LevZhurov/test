package client1

import (
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

func Client1Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}
}
