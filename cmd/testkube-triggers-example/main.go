package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.Default()
	logger.Println("Starting service on port 8080")
	http.HandleFunc("/health", healthCheckHandler)
	http.ListenAndServe(":8080", nil)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("CRASH") == "true" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Ooops, application stopped working :(")
		return
	}
	fmt.Fprint(w, "Everything is running smoothly :)")
}
