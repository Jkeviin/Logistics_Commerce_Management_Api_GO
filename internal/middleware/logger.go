package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// Logger is a middleware that logs HTTP request information
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Terminal colors
		const (
			cyan    = "\033[36m"
			green   = "\033[32m"
			yellow  = "\033[33m"
			blue    = "\033[34m"
			magenta = "\033[35m"
			red     = "\033[31m"
			reset   = "\033[0m"
			bold    = "\033[1m"
			dim     = "\033[2m"
			white   = "\033[37m"
		)

		// Create a custom response writer to capture status code
		lrw := newLoggingResponseWriter(w)

		// Execute next handler
		next.ServeHTTP(lrw, r)

		// Calculate duration
		duration := time.Since(start)

		// Format duration clearly
		var durationStr string
		switch {
		case duration >= time.Second:
			durationStr = fmt.Sprintf("%.2f seconds", duration.Seconds())
		case duration >= time.Millisecond:
			durationStr = fmt.Sprintf("%.2f milliseconds", float64(duration.Milliseconds()))
		default:
			durationStr = fmt.Sprintf("%.2f microseconds", float64(duration.Microseconds()))
		}

		// Determine status code color and message
		var statusColor, statusMsg string
		switch {
		case lrw.statusCode >= 500:
			statusColor = red
			statusMsg = "Server Error"
		case lrw.statusCode >= 400:
			statusColor = red
			statusMsg = "Client Error"
		case lrw.statusCode >= 300:
			statusColor = yellow
			statusMsg = "Redirect"
		case lrw.statusCode >= 200:
			statusColor = green
			statusMsg = "Success"
		default:
			statusColor = white
			statusMsg = "Unknown"
		}

		// Top decorative line
		fmt.Printf("%s┌─────────────────────────────────────────────────────────────┐%s\n",
			white, reset)

		// Main log with improved design
		fmt.Printf("%s│%s %s%s%s %s│%s %s%s%s %s│%s %s%d %s(%s)%s %s│%s\n",
			white, reset,
			bold+cyan, r.Method, reset, // HTTP Method
			white, reset, // Separator
			green, r.URL.Path, reset, // Path
			white, reset, // Separator
			statusColor, lrw.statusCode, dim, statusMsg, reset, // Status code with message
			white, reset, // Close
		)

		// Duration line with descriptive text
		fmt.Printf("%s│%s %sDuration:%s %s%s%s %s│%s\n",
			white, reset,
			dim, reset,
			yellow, durationStr, reset, // Duration
			white, reset, // Close
		)

		// Log relevant headers if they exist
		if auth := r.Header.Get("Authorization"); auth != "" {
			token := auth
			if len(token) > 10 {
				token = token[:10] + "..."
			}
			fmt.Printf("%s│%s %sToken:%s %s%s%s %s│%s\n",
				white, reset,
				magenta, reset,
				red, token, reset,
				white, reset, // Close
			)
		}

		// Bottom decorative line
		fmt.Printf("%s└─────────────────────────────────────────────────────────────┘%s\n",
			white, reset)
	})
}

// loggingResponseWriter is a wrapper for http.ResponseWriter that captures status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
