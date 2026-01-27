package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type LogEntry struct {
	Time       string  `json:"time"`
	Method     string  `json:"method"`
	Path       string  `json:"path"`
	RemoteAddr string  `json:"remote_addr"`
	Status     int     `json:"status"`
	DurationMs float64 `json:"duration_ms"`
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggerJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(rw, r)

		entry := LogEntry{
			Time:       time.Now().Format(time.RFC3339),
			Method:     r.Method,
			Path:       r.RequestURI,
			RemoteAddr: r.RemoteAddr,
			Status:     rw.statusCode,
			DurationMs: time.Since(start).Seconds() * 1000,
		}

		logBytes, _ := json.Marshal(entry)
		log.Println(string(logBytes))
	})
}
