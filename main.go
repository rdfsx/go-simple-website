package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	AvgGrades, Happiness float64
	Hobbies              []string
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User Name is: %s. He is %d and he has Money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := &User{"Bob", 25, -50, 4.2, 0.8,
		[]string{"football", "skate", "dance"}}
	tmpl, err := template.ParseFiles("templates/home_page.html")
	if err != nil {
		return
	}
	if tmpl == nil {
		return
	}
	err = tmpl.Execute(w, bob)
	if err != nil {
		return
	}
	fmt.Println(err)
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
