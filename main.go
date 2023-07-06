package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var templates = template.Must(template.ParseFiles("./assets/templates/base.html", "./assets/templates/body.html"))

// log handler, resources and time
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Since(start))
	})
}

// handler for root
func root() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := struct {
			Title   template.HTML
			Company string
			Slogan  string
		}{
			Title:   template.HTML("Labs example app"),
			Company: "Celonis Labs",
			Slogan:  "beyond imagination",
		}
		err := templates.ExecuteTemplate(w, "base", &b)
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

// handler public resources
func public() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./assets/public")))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", logging(root()))
	mux.Handle("/public/", logging(public()))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Println("main: running example app on port", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("main: couldn't start example app: %v\n", err)
	}
}
