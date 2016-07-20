package main

import (
	"net/http"
)

func main() {
	StartServer("8080")
}

func StartServer(port string) {
	http.HandleFunc("/hello/", makeHandler(helloHandler))

	http.ListenAndServe(":"+port, nil)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Jump jump"))
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
