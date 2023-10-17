package main

import "net/http"

func main() {

	http.HandleFunc("/home", home)
	//start your server
	http.ListenAndServe(":8080", nil) // used to start the server

}

func home(w http.ResponseWriter, h *http.Request) {

	w.Write([]byte("this is our project"))

}
