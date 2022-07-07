package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

func GetMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
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

//TODO ctx middleware

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

func PutMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
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

func DeleteMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
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
