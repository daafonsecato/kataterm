package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"gitkatas/pkg/db"
	"gitkatas/pkg/models"
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
	questionStore           *models.QuestionStore
	currentQuestionIndex    int
	totalQuestions          int
	questionStatuses        []string
	questionTrials          []int
	questionTypes           []string
	completedQuestionTrials []int
	questionTrialsLeft      []int
	lastQuestionSentIndex   int
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
	fmt.Println("Running command: ", command_to_test)

	// Run the validation task by executing the bash script
	cmd := exec.Command("/bin/bash", "-c", command_to_test)
	output, err := cmd.Output()
	if err != nil {
		errorMsg := fmt.Sprintf("Error running task: %v", err)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	fmt.Println("Output of command: ", command_to_test, " is: ", output)

	// Check if the output contains the word "success"
	if output != nil {
		fmt.Println("Output is not nil")
		fmt.Println("Output is not nil")
		fmt.Println("Output is not nil")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Accepted")
	} else {
		fmt.Println("Output is nil")
		fmt.Println("Output is nil")
		fmt.Println("Output is nil")
		fmt.Println("Output is nil")
		http.Error(w, "Output does not contain 'success'", http.StatusBadRequest)
	}
}
