package main

import "net/http"

func main() {
	// Default multiplexer
	mux := http.NewServeMux()
	// Handler to serve static files
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static",http.StripPrefix(("/static/",files)))

	// index is a handler function
	// func index(w http.ResponseWriter,r *http.Request)
	mux.HandleFunc("/",index)
	server := &http.Server {
		Addr: "0.0.0.0.8080"
		Handler: mux,
	}

	server.ListenAndServe()
}