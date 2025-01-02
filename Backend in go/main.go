package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Go! jdakjd anjksdajsdjkasn")
	})
	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Hello, %s!", vars["name"])
	})
	r.HandleFunc("/hello2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Go! 2")
	})
	http.ListenAndServe(":8080", r)
}
