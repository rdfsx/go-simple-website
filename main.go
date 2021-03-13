package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	name                 string
	age                  uint16
	money                int16
	avgGrades, happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := &User{"Bob", 25, -50, 4.2, 0.8}
	bob.setNewName("Alex")
	n, err := fmt.Fprintf(w, bob.getAllInfo())
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
