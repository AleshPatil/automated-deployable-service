package main

import (
	"fmt"
	"net/http"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).String()
	fmt.Fprintf(w, "Status: Up\nUptime: %s\n", uptime)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, Alesh!")
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	port := ":8080"
	fmt.Printf("Server starting on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
