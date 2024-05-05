package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, v interface{}, code int) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		log.Printf("error occurred while marshalling to json: %s", err)
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
	return nil
}

func respondWithError(w http.ResponseWriter, code int, errMessage string) error {
	return respondWithJSON(w, map[string]string{"error": errMessage}, code)
}
