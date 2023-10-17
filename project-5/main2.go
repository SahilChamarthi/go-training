package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", RequestIdMid(LoggingMid1(homePage1)))
	http.ListenAndServe(":8080", nil)

}

func homePage1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page invoked")
	fmt.Fprintln(w, "this is my home")
}

func LoggingMid1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqId, ok := ctx.Value(RequestIDKey).(string)
		if !ok {
			reqId = "Unknown"
		}
		log.Printf("%s : started   : %s %s ",
			reqId,
			r.Method, r.URL.Path)
		next(w, r)
	}

}
