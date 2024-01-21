package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var version = "1.0.0"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there")
	})

	http.HandleFunc("/GetVersion", GetVersionHandler)
	http.HandleFunc("/PutVersion", PutVersionHandler)

	// Set the port
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

func GetVersionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method. Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Handle GET Request
	fmt.Fprintf(w, "Current service version: %s", version)
}

func PutVersionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method. Only PUT is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Check for empty or invalid 'new_version' parameter
	newVersion := string(body)
	if newVersion == "" {
		http.Error(w, "Missing 'new_version' parameter in the request body", http.StatusBadRequest)
		return
	}

	// Update service version
	version = newVersion
	fmt.Fprintf(w, "Service version updated to: %s", version)
}
