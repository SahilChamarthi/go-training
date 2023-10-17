package main

import (
	"assignment-17/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/FetchStudent", handlers.GetStudent)
	http.ListenAndServe(":8080", nil)

}
