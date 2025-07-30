package main

import (
  "fmt"
  "net/http"
	"log"
	"html/template"
)

func main() {
	fmt.Println("Starting server...")

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// define callbacks
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/home-content", homeHandler)
	http.HandleFunc("/brand-content", brandHandler)

	// then listen and serve
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// See https://pkg.go.dev/net/http#example-HandleFunc for examples
func indexHandler(w htp.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Unable to load index.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, "Unable to load home.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func brandHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("brand.html")
	if err != nil {
		http.Error(w, "Unable to load brand.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
