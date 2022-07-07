package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

func PostMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("middleware: method not allowed")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		log.Println(r.Method, r.URL)

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		r = r.WithContext(ctx)
		next(w, r)
	}
}
