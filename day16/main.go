package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintln(w, "read body error")
				return
			}
			fmt.Fprintln(w, "Received:", string(body))
		}
	})

	http.ListenAndServe(":8080", nil)
}
