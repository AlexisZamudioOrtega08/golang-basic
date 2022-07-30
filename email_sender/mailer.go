package main

import (
	"html/template"
	"net/http"
)

type User struct {
	FirstName string
	LastName  string
	Date      string
	Adb       string
	Aca       string
	Txns      []string
	Tb        string
}

func mailSenderHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Date:      "01/01/2020",
		Aca:       "100",
		Adb:       "200",
		Tb:        "300",
	}
	t, _ := template.ParseFiles("email.html")
	t.Execute(w, user)
}

func main() {
	http.HandleFunc("/mailer", mailSenderHandler)
	// expose my server to my local network
	http.ListenAndServe(":3000", nil)
}
