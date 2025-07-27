package main

import (
  "fmt"
  "net/http"
	// "html/template"
)

func main() {
	fmt.Println("Hello")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
