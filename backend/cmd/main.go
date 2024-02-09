package main

import (
	"net/http"

	"github.com/david8128/quizard-backend/pkg/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	qc := controllers.NewQuestionController()
	qc.InitializeController()

	// Enable CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://frontend.terminal.kataterm.com:30713")
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

	// Rutas para el CRM
	r.HandleFunc("/questions", qc.GetQuestions).Methods("GET")
	r.HandleFunc("/questions/{id}", qc.GetQuestion).Methods("GET")
	r.HandleFunc("/questions", qc.CreateQuestion).Methods("POST", "OPTIONS")
	r.HandleFunc("/questions/current/{id}", qc.SetCurrentQuestion).Methods("POST", "OPTIONS")
	r.HandleFunc("/questions/{id}", qc.UpdateQuestion).Methods("PUT")
	r.HandleFunc("/questions/{id}", qc.DeleteQuestion).Methods("DELETE")
	r.HandleFunc("/question", qc.GetCurrentQuestion).Methods("GET")
	r.HandleFunc("/get_score", qc.GetScore).Methods("GET")
	r.HandleFunc("/trials", qc.GetTrials).Methods("GET")
	r.HandleFunc("/skip_question", qc.SkipQuestion).Methods("GET")
	r.HandleFunc("/submit_answer", qc.CheckMultipleChoice).Methods("POST", "OPTIONS")
	r.HandleFunc("/stage_before_actions", qc.StageBeforeActions).Methods("POST", "OPTIONS")

	// Rutas para las tareas de validación
	r.HandleFunc("/check_config", qc.CheckConfig).Methods("POST", "OPTIONS")

	// Rutas para las tareas de seed de la db
	r.HandleFunc("/dbseed", qc.DBSeed).Methods("POST", "OPTIONS")

	r.Use(corsMiddleware)

	http.ListenAndServe(":8000", r)
}
