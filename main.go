package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Go HTTP Server</title>
			<style>
				body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
				h1 { color: #00ADD8; }
				.endpoint { background: #f4f4f4; padding: 10px; margin: 10px 0; border-radius: 5px; }
			</style>
		</head>
		<body>
			<h1>🚀 Go HTTP Web Server</h1>
			<p>Welcome to the Go HTTP web service!</p>
			<h2>Available Endpoints:</h2>
			<div class="endpoint">GET <strong>/</strong> - This homepage</div>
			<div class="endpoint">GET <strong>/api/hello</strong> - Returns a JSON greeting</div>
			<div class="endpoint">GET <strong>/api/time</strong> - Returns current server time</div>
			<div class="endpoint">POST <strong>/api/echo</strong> - Echoes back your JSON request</div>
			<div class="endpoint">GET <strong>/health</strong> - Health check endpoint</div>
		</body>
		</html>
	`)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorJSON(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Message: "Hello from Go HTTP Server! 你好！",
		Time:    time.Now(),
	}
	sendJSON(w, response, http.StatusOK)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorJSON(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]interface{}{
		"current_time": time.Now().Format(time.RFC3339),
		"unix_time":    time.Now().Unix(),
		"timezone":     "UTC",
	}
	sendJSON(w, response, http.StatusOK)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorJSON(w, "Method not allowed, use POST", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		sendErrorJSON(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"received": data,
		"time":     time.Now().Format(time.RFC3339),
	}
	sendJSON(w, response, http.StatusOK)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorJSON(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	}
	sendJSON(w, response, http.StatusOK)
}

func sendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func sendErrorJSON(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
		log.Printf("Completed in %v", time.Since(start))
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", loggingMiddleware(homeHandler))
	mux.HandleFunc("/api/hello", loggingMiddleware(helloHandler))
	mux.HandleFunc("/api/time", loggingMiddleware(timeHandler))
	mux.HandleFunc("/api/echo", loggingMiddleware(echoHandler))
	mux.HandleFunc("/health", loggingMiddleware(healthHandler))

	port := "8080"
	addr := fmt.Sprintf(":%s", port)

	log.Printf("🚀 Starting HTTP server on http://localhost%s", addr)
	log.Printf("📝 Press Ctrl+C to stop the server")

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
