package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
	})
}

func main() {
	http.HandleFunc("/health", healthHandler)

	port := "8080"

	fmt.Printf("Service started on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
