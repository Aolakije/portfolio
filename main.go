package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Parse templates at startup (fail fast if broken)
	tmpl := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

	mux := http.NewServeMux()

	// Serve static files from ./static at /static/
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Application routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Printf("template execute error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	log.Print("Starting new server on :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
