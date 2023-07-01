package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

}

func formHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting serve at port: 8000\n")
	if err := http.ListenAndServe("8000", nil); err != nil {
		log.Fatal(err)
	}
}
