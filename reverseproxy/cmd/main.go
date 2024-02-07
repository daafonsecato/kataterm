package main

import (
	"log"
	"net/http"

	"github.com/daafonsecato/kataterm-reverseproxy/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Enable CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://terminal.kataterm.com")
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
	// Session controller
	sc := handlers.NewSessionController()

	// Start session controller in a separate goroutine
	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/create", sc.CreateKubernetesPodHandler).Methods("GET")
		r.HandleFunc("/terminate", sc.TerminateMachineHandler).Methods("POST", "OPTIONS")
		r.HandleFunc("/terminatemults", sc.TerminateMultipleMachinesHandler).Methods("POST", "OPTIONS")
		r.Use(corsMiddleware)
		log.Fatal(http.ListenAndServe(":9090", r))
	}()

	proxy := http.HandlerFunc(sc.ReverseProxyHandler)
	// Start reverse proxy in a separate goroutine
	go func() {

		log.Fatal(http.ListenAndServe(":7070", corsMiddleware(proxy)))
	}()

	// Wait indefinitely to keep the main goroutine running
	select {}
}
