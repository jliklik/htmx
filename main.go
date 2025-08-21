package main

import (
  "fmt"
  "net/http"
	"log"
	"html/template"
	"sync"
)

type Server struct {
	counter int
	mu      sync.Mutex
}

func main() {

	s := &Server{}
	fmt.Println("Starting server...")

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// define callbacks
	http.HandleFunc("/", s.indexHandler)

	http.HandleFunc("/home-content", s.homeHandler)
	http.HandleFunc("/brand-content", s.brandHandler)
	http.HandleFunc("/service-content", s.serviceHandler)
	http.HandleFunc("/booking-content", s.bookingHandler)
	http.HandleFunc("/review-content", s.reviewHandler)
	http.HandleFunc("/slide", s.homeSlideHandler)

	// then listen and serve
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// See https://pkg.go.dev/net/http#example-HandleFunc for examples
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Unable to load index.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, "Unable to load home.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func (s *Server) brandHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("brand.html")
	if err != nil {
		http.Error(w, "Unable to load brand.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func (s *Server) serviceHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("service.html")
	if err != nil {
		http.Error(w, "Unable to load service.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}


func (s *Server) bookingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("booking.html")
	if err != nil {
		http.Error(w, "Unable to load booking.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func (s *Server) reviewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("review.html")
	if err != nil {
		http.Error(w, "Unable to load review.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func (s *Server) homeSlideHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	if (s.counter >= 3) {
		s.counter = 0
	}
	fmt.Fprintf(w,
	  `
		<div id="home-slide" class="slide-fade overflow-hidden flex flex-col flex-1 bg-black" hx-get="/slide" hx-swap="outerHTML swap:1s settle:1s" hx-trigger="every 3s" hx-target="#home-slide">
			<img src='/static/shop_%d.jpg' class="w-full h-full object-cover">
		</div>
		`, 
	 s.counter)
}
