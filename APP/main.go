package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "USER APP : HEALTH IS OK")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hitting Root Handler")
}

func main() {
	fmt.Println("User APP is getting initialized")
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
