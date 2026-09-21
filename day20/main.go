package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Alice", Age: 20},
	{ID: 2, Name: "Bob", Age: 21},
	{ID: 3, Name: "Charlie", Age: 22},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			getUsers(w, r)
		case http.MethodPost:
			createUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

		return
	}

	if strings.HasPrefix(r.URL.Path, "/users/") {
		idText := strings.TrimPrefix(r.URL.Path, "/users/")

		id, err := strconv.Atoi(idText)
		if err != nil {
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			getUser(w, r, id)
		case http.MethodDelete:
			deleteUser(w, r, id)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	http.NotFound(w, r)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request, id int) {
	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "user not found", http.StatusNotFound)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "user not found", http.StatusNotFound)

}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("请求：", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})

}

func main() {
	handler := logger(http.HandlerFunc(userHandler))

	http.Handle("/", handler) //收到HTTP请求之后把它交给这个handler处理

	fmt.Println("serve started at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("serve error:", err)
	}
}
