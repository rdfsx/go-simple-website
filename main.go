package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello Go")
	fmt.Println(n, err)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Number +")
	fmt.Println(n, err)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
