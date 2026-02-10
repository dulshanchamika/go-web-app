package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// staticRoot picks a static directory that works both locally and in containers.
// - If "/static" exists (common in Docker/k8s), use it.
// - Otherwise use "./static" (local dev / tests).
func staticRoot() string {
	if info, err := os.Stat("/static"); err == nil && info.IsDir() {
		return "/static"
	}
	return "static"
}

func serve(root, file string) http.HandlerFunc {
	fullPath := filepath.Join(root, file)
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fullPath)
	}
}

func routes() http.Handler {
	root := staticRoot()
	mux := http.NewServeMux()

	// Static assets (css/js/images) under /static/...
	mux.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(root))),
	)

	// Pages
	mux.HandleFunc("/", serve(root, "home.html"))
	mux.HandleFunc("/home", serve(root, "home.html"))
	mux.HandleFunc("/courses", serve(root, "courses.html"))
	mux.HandleFunc("/about", serve(root, "about.html"))
	mux.HandleFunc("/contact", serve(root, "contact.html"))

	return mux
}

func main() {
	log.Println("Listening on :8080")

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: routes(),
	}

	log.Fatal(srv.ListenAndServe())
}
