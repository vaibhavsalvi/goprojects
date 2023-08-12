package main

import (
	"fmt"
	"net/http"

	"helloworld/service"
)

func handler(w http.ResponseWriter, r *http.Request) {
	userData := getUser()
	fmt.Fprintf(w, "Hello, Web!"+userData.Username)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func getUser() service.User {
	userData := service.User{
		ID:       1,
		Username: "vaibhav",
		Email:    "v.g",
	}
	return userData
}
