package slog

import (
	"log"
	"net/http"
	"os"
)

// API请求日志记录
func ApiLog() {
	file, err := os.OpenFile("../../log/api_requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "REQUEST: ", log.Ldate|log.Ltime|log.LUTC)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		w.Write([]byte("Hello, world!"))
	})

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
