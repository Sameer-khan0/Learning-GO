package main

import (
	"backend-in-go/mypackage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{num1}/{num2}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		headers := r.Header.Get("Authorization")
		// bodydata := r.Body.name
		fmt.Println(headers)
		// fmt.Println(bodydata)
		num1, err1 := strconv.Atoi(vars["num1"])
		num2, err2 := strconv.Atoi(vars["num2"])
		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid numbers provided in the URL", http.StatusBadRequest)
			return
		}
		result := mypackage.Add(num1, num2)
		fmt.Fprintf(w, "Result of %d and %d is %d", num1, num2, result)
	})

	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Hello, %s! you are great!", vars["name"])
		fmt.Fprintf(w, "Hello, %s! you are great!", vars["name"])
	})

	r.HandleFunc("/hello2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Go! 2")
	})
	http.ListenAndServe(":8080", r)
}
