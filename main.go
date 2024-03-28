package main

import (
	"fmt"
	"log"
	"net/http"
)

func checkMethod(w http.ResponseWriter, r *http.Request, allowedMethod string) bool {
	if r.Method != allowedMethod {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func checkPath(w http.ResponseWriter, r *http.Request, expectedPath string) bool {
	if r.URL.Path != expectedPath {
		http.Error(w, "404 not found", http.StatusNotFound)
		return false
	}
	return true
}



func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, "GET") {
		return
	}

	if !checkPath(w, r, "/about") {
		return
	}

	http.ServeFile(w, r, "./static/about.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, "GET") {
		return
	}

	if !checkPath(w, r, "/contact") {
		return
	}

	http.ServeFile(w, r, "./static/contact.html")
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, "GET") {
		return
	}

	if !checkPath(w, r, "/services") {
		return
	}

	http.ServeFile(w, r, "./static/services.html")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/services", servicesHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
