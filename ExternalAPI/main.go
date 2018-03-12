package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hitting Root Handler")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "EXTERNAL API : HEALTH IS OK")
}

func main() {
	fmt.Println("External API is getting initialized")
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":7000", nil))
}
