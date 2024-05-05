package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"
)

type requestParams struct {
	Body string `json:"body"`
}

func handleValidateChirp(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	reqParams := requestParams{}
	err := decoder.Decode(&reqParams)
	if err != nil {
		log.Printf("Error decoding parameters, %s", err)
		respondWithError(w, 500, "Something went wrong")
		return
	}
	if len(reqParams.Body) > 140 {
		log.Printf("Chirp is too long")
		respondWithError(w, 400, "Chirp is too long")
	} else {
		log.Printf("Valid chirp")
		cleanBody := replaceProfaneWords(reqParams.Body)
		respondWithJSON(w, map[string]string{"cleaned_body": cleanBody}, 200)
	}
}

func replaceProfaneWords(body string) string {
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}

	words := strings.Split(body, " ")
	cleanedWords := make([]string, 0)

	for _, word := range words {
		cleanedWord := word
		if slices.Contains(profaneWords, strings.ToLower(word)) {
			cleanedWord = "****"
		}
		cleanedWords = append(cleanedWords, cleanedWord)
	}
	cleanBody := strings.Join(cleanedWords, " ")
	return cleanBody
}
