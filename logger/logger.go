package logger

import (
	"log"
	"net/http"
	"time"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		s := &StatusWriter{ResponseWriter: w,}
		targetMux.ServeHTTP(s, r)

		log.Printf(
			"%s %s %s %d %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			s.Status,
			time.Since(start),
		)
	})

}
