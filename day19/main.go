package main

import (
	"fmt"
	"net/http"
	"strings"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, `[
			{"id":1, "name":"Alice", "age":20},
			{"id":2, "name":"Bob", "age":22}
		]`)

		return
	}

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintln(w, `{"id":3, "name":"Charlie", "age":21}`)

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w, `{"error":"method not allowed"}`)
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)

	id := strings.TrimPrefix(r.URL.Path, "/users/")

	fmt.Println("User ID:", id)

	if r.Method == "GET" {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, `{"id":1,"name":"Alice","age":20}`)

		return
	}

	if r.Method == "PUT" {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, `{"id":1, "name":"Alice Updated", "age":21}`)

		return
	}

	if r.Method == "DELETE" {

		w.WriteHeader(http.StatusNoContent)

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w, `{"error":"method not allowed"}`)
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/1", userHandler)

	fmt.Println("server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
