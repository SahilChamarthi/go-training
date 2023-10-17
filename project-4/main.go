package main

import (
	"net/http"
	"project-4/handlers"
)

func main() {

	http.HandleFunc("/home", handlers.Home)
	//starting server
	panic(http.ListenAndServe(":8080", nil))

}
