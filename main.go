package main

import (
	"log"
	"net/http"
)

func serve(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))
	// Serve pages
	http.HandleFunc("/", serve("/static/home.html")) // ✅ root works now
	http.HandleFunc("/home", serve("/static/home.html"))
	http.HandleFunc("/courses", serve("/static/courses.html"))
	http.HandleFunc("/about", serve("/static/about.html"))
	http.HandleFunc("/contact", serve("/static/contact.html"))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
