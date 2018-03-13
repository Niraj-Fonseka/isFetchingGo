package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hitting Root Handler")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "EXTERNAL API : HEALTH IS OK")
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	fmt.Println("Request Recived to External API")
	fmt.Fprintf(w, "PAYLOAD FROM EXTERNAL API")
}

func main() {
	fmt.Println("External API is getting initialized")
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/data", getDataHandler)
	http.HandleFunc("/", rootHandler)
	log.Println("Running in port 7000")
	log.Fatal(http.ListenAndServe(":7000", nil))
}
