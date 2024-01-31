package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"validator/pkg/db"
	"validator/pkg/models"
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

func (controller *QuestionController) CheckConfig(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("Command to run:", command_to_test)
	// Run the validation task by executing the bash script
	cmd := exec.Command("/bin/bash", "-c", command_to_test)
	output, err := cmd.Output()
	if err != nil {
		errorMsg := fmt.Sprintf("Error running task: %v", command_to_test)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}
	fmt.Println("Output:", string(output))
	// Check if the output contains the word "success"
	if strings.Contains(string(output), "success") {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Accepted"))
	} else {
		errormsg := fmt.Sprintf("Error: %v", string(output))
		http.Error(w, errormsg, http.StatusBadRequest)
	}
}
