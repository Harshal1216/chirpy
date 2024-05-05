package main

import (
	"log"
	"net/http"
)

func main() {
	filepathRoot := "."
	port := "8080"
	filePath := http.Dir(filepathRoot)
	homePageHandler := http.StripPrefix("/app", http.FileServer(filePath))
	config := apiConfig{
		fileServerHits: 0,
	}

	serveMux := http.NewServeMux()
	serveMux.Handle("/app/*", config.middlewareMetricsInc(homePageHandler))
	serveMux.HandleFunc("GET /api/healthz", handlerReadiness)
	serveMux.HandleFunc("GET /admin/metrics", config.handlerCount)
	serveMux.HandleFunc("POST /api/validate_chirp", handleValidateChirp)
	serveMux.HandleFunc("/api/reset", config.handlerReset)

	server := http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
	log.Printf("Serving files from %s on port: %s", filepathRoot, port)
	log.Fatal(server.ListenAndServe())

}
