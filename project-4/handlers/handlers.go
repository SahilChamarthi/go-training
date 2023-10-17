package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, h *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(User{"sahil", []string{"cricket", "swim"}})

	if err != nil {

		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, http.StatusText(http.StatusInternalServerError))
		return
		//w.Write([]byte("somthing wenr wrong"))

	}

	w.WriteHeader(http.StatusOK)

	w.Write(jsonData)

}

type User struct {
	UserID  string
	Hobbies []string
}
