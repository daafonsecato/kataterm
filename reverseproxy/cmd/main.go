package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daafonsecato/kataterm-reverseproxy/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Enable CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://frontend.terminal.kataterm.com:30713")
			w.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, My-Service, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			// Handle OPTIONS request
			if r.Method == "OPTIONS" || r.Method == "HEAD" {
				fmt.Println("OPTIONS request")
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
	// Session controller
	sc := handlers.NewSessionController()

	r := mux.NewRouter()
	r.HandleFunc("/create", sc.CreateKubernetesPodHandler).Methods("GET", "OPTIONS", "HEAD")
	r.HandleFunc("/terminate", sc.TerminateMachineHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/terminatemults", sc.TerminateMultipleMachinesHandler).Methods("POST", "OPTIONS")
	r.Use(corsMiddleware)
	log.Fatal(http.ListenAndServe(":9090", r))

}
