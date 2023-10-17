package handlers

import (
	"assignment-17/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetStudent(w http.ResponseWriter, r *http.Request) {

	studIdString := r.URL.Query().Get("student_id")

	studId, err := strconv.ParseUint(studIdString, 10, 64)

	if err != nil {

		log.Println("Error: ", err)

		errorInConversion := map[string]string{"msg": "not a valid number"}

		jsonData, err := json.Marshal(errorInConversion)

		if err != nil {
			log.Println("Error while converting error to json", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		w.WriteHeader(http.StatusBadRequest)

		w.Write(jsonData)

		return

	}

	sData, err := models.FetchStudent(studId)

	if err != nil {

		fetchError := map[string]string{"msg": "user not found"}

		errData, err := json.Marshal(fetchError)

		if err != nil {

			log.Println("Error while parsing fetchuser error conversion: ", err)

			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

			return

		}

		w.WriteHeader(http.StatusInternalServerError)

		w.Write(errData)

		return

	}

	studData, err := json.Marshal(sData)

	if err != nil {

		log.Println("Error while converting user data to json", err)

		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

		return

	}

	w.Write(studData)

}
