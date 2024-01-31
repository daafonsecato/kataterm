package main

import (
	"net/http"

	"gitkatas/pkg/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	qc := controllers.NewQuestionController()

	// Enable CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://terminal.kataterm.com,http://terminal.kataterm.com:8000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			// Handle OPTIONS request
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	r.HandleFunc("/stage_before_actions", qc.StageBeforeActions).Methods("POST", "OPTIONS")

	r.Use(corsMiddleware)

	http.ListenAndServe(":8095", r)
}
