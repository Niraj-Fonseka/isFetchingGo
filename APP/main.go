package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "USER APP : HEALTH IS OK")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hitting Root Handler")
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getDataHandler is being called")
	fmt.Println("IsFetching Should Set")
	fmt.Printf("IsFetching : %t \n", isFetching)
	if isFetching {
		fmt.Println("Is fetch is true ")
		data, _ := <-dataChan
		fmt.Printf("Data Reciveved ! \n --- >PAYLOAD : %s \n\n", data)
		fmt.Fprintf(w, data)
	} else {
		fmt.Printf("Is fetch is not true , aka first call")
		fmt.Printf("isFetching Value before SET : %t \n", isFetching)
		isFetching = true
		fmt.Println("Set is fetching to true")
		fmt.Printf("isFetching Value after SET : %t \n", isFetching)
		data := getDataFromExternal()
		fmt.Printf("Data Reciveved ! \n --- >PAYLOAD : %s \n\n", data)
		dataBefore, err := <-dataChan
		fmt.Printf("Error Status : %t \n", err)
		fmt.Printf("Data in the channel before insert: %s \n", dataBefore)
		dataChan <- data
		fmt.Fprintf(w, data)
	}
}

func getDataFromExternal() string {
	fmt.Println("getDataFromExternal function getting called ")
	response, err := http.Get("http://localhost:7000/data")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("Payload inside getDataFromExternal %s \n", contents)
		dataChan <- string(contents)
		fmt.Println("Done Pushing Data into the dataChan channel")
		return string(contents)
	}
	return "null"
}

var isFetching bool
var dataChan = make(chan string)

func main() {
	fmt.Println("User APP is getting initialized")
	isFetching = false

	fmt.Printf("IsFetching : Init : %t \n", isFetching)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/getData", getDataHandler)
	http.HandleFunc("/", rootHandler)
	log.Println("Running User APP in port :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
