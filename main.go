package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Global variables to store our redirects and redirect status code
var (
	redirects      sync.Map
	redirectStatus int
)

// init function runs before main, used to set up our redirects and redirect status
func init() {
	// Set up redirect status
	redirectStatus = http.StatusMovedPermanently // Default to 301 Moved Permanently
	if status := os.Getenv("REDIRECT_STATUS"); status != "" {
		if code, err := strconv.Atoi(status); err == nil && code >= 300 && code < 400 {
			redirectStatus = code
		} else {
			log.Printf("Invalid REDIRECT_STATUS '%s', using default 301", status)
		}
	}

	// Set up redirects
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "REDIRECT_") && env != "REDIRECT_STATUS" {
			parts := strings.SplitN(env, "=", 2)
			if len(parts) == 2 {
				key := "/" + strings.ToLower(strings.TrimPrefix(parts[0], "REDIRECT_"))
				redirects.Store(key, parts[1])
			}
		}
	}
}

// handleRedirect is our main request handler
func handleRedirect(w http.ResponseWriter, r *http.Request) {
	path := strings.ToLower(r.URL.Path)

	// Check if we have a redirect for this path
	if target, ok := redirects.Load(path); ok {
		http.Redirect(w, r, target.(string), redirectStatus)
		return
	}

	// If no redirect found, return 404
	http.NotFound(w, r)
}

func main() {
	// Set up our single route
	http.HandleFunc("/", handleRedirect)

	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	addr := ":" + port
	log.Printf("Server starting on %s", addr)
	log.Printf("Redirect status: %d", redirectStatus)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
