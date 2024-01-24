package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/david8128/quizard-backend/pkg/db"
	"github.com/david8128/quizard-backend/pkg/models"
	"github.com/gorilla/mux"
)

type Question struct {
	ID                 string          `json:"ID"`
	Content_Text       string          `json:"Content_Text"`
	Hint               string          `json:"hint"`
	Subtext            string          `json:"subtext"`
	Type_Question      string          `json:"type_question"`
	Staging_Message    string          `json:"staging_message"`
	Options            json.RawMessage `json:"options"`
	Before_Actions     json.RawMessage `json:"before_actions"`
	Answer             string          `json:"answer"`
	Test_spec_filename string          `json:"test_spec_filename"`
	Trials             string          `json:"trials"`
}

type QuestionStore struct {
	db *sql.DB
}

type QuestionController struct {
	questionStore        *models.QuestionStore
	currentQuestionIndex int
	totalQuestions       int
	questionStatuses     []string
}

func NewQuestionController() *QuestionController {
	db, err := db.InitDB()
	if err != nil {
		panic("Error initializing DB")
	}

	questionStore := models.NewQuestionStore(db)

	return &QuestionController{
		questionStore:        questionStore,
		currentQuestionIndex: 0,
		totalQuestions:       0,
		questionStatuses:     []string{},
	}
}

func (controller *QuestionController) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := controller.questionStore.GetQuestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(questions) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (controller *QuestionController) CheckMultipleChoice(w http.ResponseWriter, r *http.Request) {
	// Obtener el payload del request
	var payload struct {
		Answer string `json:"answer"`
	}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener la respuesta actual
	currentQuestion, err := controller.questionStore.GetQuestion(strconv.Itoa(controller.currentQuestionIndex))
	if err != nil {
		http.Error(w, "No hay pregunta actual", http.StatusBadRequest)
		return
	}

	// Comparar la respuesta del payload con la respuesta actual
	if payload.Answer == currentQuestion.Answer {
		controller.SetCurrentQuestionIndex(controller.currentQuestionIndex + 1)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (controller *QuestionController) GetCurrentQuestion(w http.ResponseWriter, r *http.Request) {

	id := strconv.Itoa(controller.currentQuestionIndex)
	question, err := controller.questionStore.GetQuestion(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	controller.GetTotalQuestions()
	questionDetails := controller.currentCuestionDetails(question.TypeQuestion, question)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questionDetails)
}

func (controller *QuestionController) currentCuestionDetails(typeQuestion string, question *models.Question) interface{} {

	if question.TypeQuestion != "config_test" {
		questionDetails := struct {
			ID                    string                   `json:"id"`
			Text                  string                   `json:"text"`
			Hint                  string                   `json:"hint"`
			Subtext               string                   `json:"subtext"`
			Type                  string                   `json:"type"`
			StagingMessage        string                   `json:"staging_message"`
			Options               []string                 `json:"options"`
			Before_Actions        []map[string]interface{} `json:"before_actions"`
			Answer                string                   `json:"answer"`
			Trials                string                   `json:"trials"`
			TotalQuestions        int                      `json:"total_questions"`
			CurrentQuestionNumber int                      `json:"current_question_number"`
			AnswerStatuses        []string                 `json:"answer_statuses"`
		}{
			ID:                    question.ID,
			Text:                  question.ContentText,
			Hint:                  question.Hint,
			Subtext:               question.Subtext,
			Type:                  question.TypeQuestion,
			StagingMessage:        question.StagingMessage,
			Options:               parseOptions(question.Options),
			Before_Actions:        parseBeforeActions(question.BeforeActions),
			Answer:                question.Answer,
			Trials:                question.Trials,
			TotalQuestions:        controller.totalQuestions,
			CurrentQuestionNumber: controller.currentQuestionIndex,
			AnswerStatuses:        controller.questionStatuses,
		}
		return questionDetails
	} else {
		questionDetails := struct {
			ID                    string                   `json:"id"`
			Text                  string                   `json:"text"`
			Hint                  string                   `json:"hint"`
			Subtext               string                   `json:"subtext"`
			Type                  string                   `json:"type"`
			StagingMessage        string                   `json:"staging_message"`
			Test_spec_filename    string                   `json:"test_spec_filename"`
			Before_Actions        []map[string]interface{} `json:"before_actions"`
			Trials                string                   `json:"trials"`
			TotalQuestions        int                      `json:"total_questions"`
			CurrentQuestionNumber int                      `json:"current_question_number"`
			AnswerStatuses        []string                 `json:"answer_statuses"`
		}{
			ID:                    question.ID,
			Text:                  question.ContentText,
			Hint:                  question.Hint,
			Subtext:               question.Subtext,
			Type:                  question.TypeQuestion,
			StagingMessage:        question.StagingMessage,
			Before_Actions:        parseBeforeActions(question.BeforeActions),
			Test_spec_filename:    question.TestSpecFilename,
			Trials:                question.Trials,
			TotalQuestions:        controller.totalQuestions,
			CurrentQuestionNumber: controller.currentQuestionIndex,
			AnswerStatuses:        controller.questionStatuses,
		}
		return questionDetails
	}
}

func parseOptions(options json.RawMessage) []string {
	var parsedOptions []string
	json.Unmarshal(options, &parsedOptions)
	return parsedOptions
}

func parseBeforeActions(beforeActions json.RawMessage) []map[string]interface{} {
	var parsedBeforeActions []map[string]interface{}
	json.Unmarshal(beforeActions, &parsedBeforeActions)
	return parsedBeforeActions
}

// SetCurrentQuestionByID sets the current question by ID.
func (controller *QuestionController) SetCurrentQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	controller.SetCurrentQuestionIndex(idInt)
}

// SetCurrentQuestionIndex sets the current question index.
func (controller *QuestionController) SetCurrentQuestionIndex(index int) {
	controller.currentQuestionIndex = index
}

// GetCurrentQuestionIndex retrieves the current question index.
func (controller *QuestionController) GetCurrentQuestionIndex() int {
	return controller.currentQuestionIndex
}

// GetTotalQuestions calculates the total number of questions for a session.
func (controller *QuestionController) GetTotalQuestions() int {
	num_questions, err := controller.questionStore.GetNumberOfQuestions()
	fmt.Printf("Number of questions: %d\n", num_questions)
	if err != nil {
		controller.totalQuestions = -1
	}
	controller.totalQuestions = num_questions
	return controller.totalQuestions
}

// SetQuestionStatus sets the status of a question at a given index.
func (controller *QuestionController) SetQuestionStatus(index int, status string) {
	if index >= 0 && index < len(controller.questionStatuses) {
		controller.questionStatuses[index] = status
	}
}

// GetTotalQuestionsStatus retrieves the statuses of all questions.
func (controller *QuestionController) GetTotalQuestionsStatus() []string {
	return controller.questionStatuses
}

func (controller *QuestionController) GetQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	question, err := controller.questionStore.GetQuestion(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func (controller *QuestionController) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var question Question
	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if question.Type_Question == "config_test" {
		if question.Test_spec_filename == "" {
			http.Error(w, "test_spec_filename is required", http.StatusBadRequest)
			return
		}
	} else {
		if question.Options == nil || question.Answer == "" {
			http.Error(w, "options and answer are required", http.StatusBadRequest)
			return
		}
	}

	createdQuestion, err := controller.questionStore.CreateQuestion(question.ID, question.Content_Text, question.Hint, question.Subtext, question.Type_Question, question.Staging_Message, question.Options, question.Before_Actions, question.Answer, question.Test_spec_filename, question.Trials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdQuestion)
	w.WriteHeader(http.StatusCreated)
}

func (controller *QuestionController) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedQuestion Question
	err := json.NewDecoder(r.Body).Decode(&updatedQuestion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedQuestion.Type_Question == "config_test" {
		if updatedQuestion.Test_spec_filename == "" {
			http.Error(w, "test_spec_filename is required", http.StatusBadRequest)
			return
		}
	} else {
		if updatedQuestion.Options == nil || updatedQuestion.Answer == "" {
			http.Error(w, "options and answer are required", http.StatusBadRequest)
			return
		}
	}

	err = controller.questionStore.UpdateQuestion(id, updatedQuestion.Content_Text, updatedQuestion.Hint, updatedQuestion.Subtext, updatedQuestion.Type_Question, updatedQuestion.Staging_Message, updatedQuestion.Options, updatedQuestion.Before_Actions, updatedQuestion.Answer, updatedQuestion.Test_spec_filename, updatedQuestion.Trials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // Send 200 response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedQuestion)
}

func (controller *QuestionController) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := controller.questionStore.DeleteQuestion(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type Task struct {
	Name   string
	TaskID string
}

func NewTask(name string, taskid string) *Task {
	return &Task{
		Name:   name,
		TaskID: taskid,
	}
}

func (controller *QuestionController) CheckConfig(w http.ResponseWriter, r *http.Request) {
	// Add your configuration checking logic here
	// Return an error if the configuration is invalid

	// Check if the script exists for the given task ID
	var data struct {
		ID int `json:"ID"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	question, err := controller.questionStore.GetQuestion(strconv.Itoa(data.ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	command_to_test := question.TestSpecFilename

	// Run the validation task by executing the bash script
	cmd := exec.Command("/bin/bash","-c",  command_to_test)
	output, err := cmd.Output()
	if err != nil {
		errorMsg := fmt.Sprintf("Error running task: %v", command_to_test)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	// Check if the output contains the word "success"
	if strings.Contains(string(output), "success") {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Accepted"))

		controller.SetCurrentQuestionIndex(controller.currentQuestionIndex + 1)
	} else {
		http.Error(w, "Output does not contain 'success'", http.StatusBadRequest)
	}
}

func (controller *QuestionController) StageBeforeActions(w http.ResponseWriter, r *http.Request) {
	// Add your configuration checking logic here
	// Return an error if the configuration is invalid

	// Check if the script exists for the given task ID
	type BeforeAction struct {
		Type    string `json:"type"`
		Command string `json:"command"`
		Shell   bool   `json:"shell"`
	}

	var data struct {
		ID            string         `json:"ID"`
		BeforeActions []BeforeAction `json:"before_actions"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	

	question, err := controller.questionStore.GetQuestion(data.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	beforeActions := make([]map[string]interface{}, 0)
	err = json.Unmarshal(question.BeforeActions, &beforeActions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	command_to_test := beforeActions[0]["command"].(string)

	// Run the validation task by executing the bash script
	cmd := exec.Command("/bin/bash", "-c", command_to_test)
	output, err := cmd.Output()
	if err != nil {
		errorMsg := fmt.Sprintf("Error running task: %v", command_to_test)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	questionDetails := controller.currentCuestionDetails(question.TypeQuestion, question)

	// Check if the output contains the word "success"
	if output != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(questionDetails)
	} else {
		http.Error(w, "Output does not contain 'success'", http.StatusBadRequest)
	}
}

func (controller *QuestionController) DBSeed(w http.ResponseWriter, r *http.Request) {
	var questions []Question
	err := json.NewDecoder(r.Body).Decode(&questions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, question := range questions {
		if question.Type_Question == "config_test" {
			if question.Test_spec_filename == "" {
				http.Error(w, "test_spec_filename is required", http.StatusBadRequest)
				return
			}
		} else {
			if question.Options == nil || question.Answer == "" {
				http.Error(w, "options and answer are required", http.StatusBadRequest)
				return
			}
		}

		_, err := controller.questionStore.CreateQuestion(question.ID, question.Content_Text, question.Hint, question.Subtext, question.Type_Question, question.Staging_Message, question.Options, question.Before_Actions, question.Answer, question.Test_spec_filename, question.Trials)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
