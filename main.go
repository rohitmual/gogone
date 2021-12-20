package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to Go Web Dev!!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the Server on port: 3000")
	http.ListenAndServe(":3000", nil)
}
