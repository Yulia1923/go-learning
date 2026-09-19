package main

import (
	"fmt"
	"net/http"
	"strings"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")

	name, ok := users[id]

	if !ok {
		fmt.Fprintln(w, "User not found")
		return
	}

	fmt.Fprintln(w, name)
}

var users = map[string]string{
	"1": "Alice",
	"2": "Bob",
	"3": "Charlie",
}

func main() {
	http.HandleFunc("/user/", userHandler)

	http.ListenAndServe(":8080", nil)
}
