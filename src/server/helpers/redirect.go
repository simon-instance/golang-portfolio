package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithJSON sends back a json response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	log.SetPrefix("[redirect.jsonResponseError] :: ")
	response, err := json.Marshal(payload)

	if err != nil {
		RespondWithError(w, code, err.Error())
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Fatalf("Error => %v", err)
	}
}

// RespondWithError sends back json response as an error
func RespondWithError(w http.ResponseWriter, code int, msg string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := json.Marshal(map[string]string{"message": msg})
		if err != nil {
			log.Fatalf("marshalError: %v", err.Error())
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		_, err = w.Write(response)
		if err != nil {
			log.Fatalf("Error => %v", err)
		}
	})
}
