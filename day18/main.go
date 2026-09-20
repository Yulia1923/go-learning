package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:   2,
		Name: "Bob",
		Age:  20,
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	w.Write(data)
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}
