package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

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
func (controller *QuestionController) InitializeController() {
	questions, err := controller.questionStore.GetQuestions()
	if err != nil {
		panic(err)
	}
	controller.SetCurrentQuestionIndex(1)
	controller.lastQuestionSentIndex = 0
	controller.totalQuestions = len(questions)
	controller.questionStatuses = make([]string, controller.totalQuestions)
	controller.questionTrials = make([]int, controller.totalQuestions)
	controller.questionTypes = make([]string, controller.totalQuestions)
	controller.completedQuestionTrials = make([]int, controller.totalQuestions)
	controller.questionTrialsLeft = make([]int, controller.totalQuestions)

	for _, question := range questions {
		questionID, err := strconv.Atoi(question.ID)
		if err != nil {
			panic(err)
		}
		questionID = questionID - 1
		controller.questionTypes[questionID] = question.TypeQuestion
		trials, err := strconv.Atoi(question.Trials)
		if err != nil {
			panic(err)
		}
		controller.completedQuestionTrials[questionID] = 0
		controller.questionTrialsLeft[questionID] = trials
		controller.questionTrials[questionID] = trials
	}

	controller.SetCurrentQuestionIndex(1)
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
		index, err := strconv.Atoi(currentQuestion.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		controller.questionTrialsLeft[controller.currentQuestionIndex-1] = controller.questionTrialsLeft[controller.currentQuestionIndex-1]
		controller.completedQuestionTrials[controller.currentQuestionIndex-1] = controller.questionTrialsLeft[controller.currentQuestionIndex-1]
		controller.SetCurrentQuestionIndex(index + 1)
		w.WriteHeader(http.StatusOK)
	} else {
		controller.questionTrialsLeft[controller.currentQuestionIndex-1] = controller.questionTrialsLeft[controller.currentQuestionIndex-1] - 1
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
	trials, err := strconv.Atoi(question.Trials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if controller.lastQuestionSentIndex == (controller.currentQuestionIndex - 1) {
		controller.questionTrialsLeft[controller.currentQuestionIndex-1] = trials
	} else {
		controller.questionTrialsLeft[controller.currentQuestionIndex-1] = trials - 1
	}
	controller.lastQuestionSentIndex = controller.currentQuestionIndex
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
			TrialsLeft            int                      `json:"trials_left"`
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
			TrialsLeft:            controller.questionTrialsLeft[controller.currentQuestionIndex-1],
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
			TrialsLeft            int                      `json:"trials_left"`
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
			TrialsLeft:            controller.questionTrialsLeft[controller.currentQuestionIndex-1],
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

// GetScore calculates the score for config_test questions and multiple_choice questions.
func (controller *QuestionController) GetScore(w http.ResponseWriter, r *http.Request) {
	globalCompletedSum := sum(controller.completedQuestionTrials)
	maxCompletedScore := sum(controller.questionTrials)
	globalScore := float64(globalCompletedSum) / float64(maxCompletedScore) * 100

	configTestCompletedSum := 0
	configTestMaxScore := 0
	multipleChoiceCompletedSum := 0
	multipleChoiceMaxScore := 0

	for i, questionType := range controller.questionTypes {
		if questionType == "config_test" {
			configTestCompletedSum += controller.completedQuestionTrials[i]
			configTestMaxScore += controller.questionTrials[i]
		} else if questionType == "multiple_choice" {
			multipleChoiceCompletedSum += controller.completedQuestionTrials[i]
			multipleChoiceMaxScore += controller.questionTrials[i]
		}
	}

	configTestScore := float64(configTestCompletedSum) / float64(configTestMaxScore) * 100
	multipleChoiceScore := float64(multipleChoiceCompletedSum) / float64(multipleChoiceMaxScore) * 100

	response := map[string]float64{
		"Global_Score":          globalScore,
		"Config_Test_Score":     configTestScore,
		"Multiple_Choice_Score": multipleChoiceScore,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (controller *QuestionController) GetTrials(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Trials: %x", controller.questionTrialsLeft)
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
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

func (controller *QuestionController) SkipQuestion(w http.ResponseWriter, r *http.Request) {
	controller.SetCurrentQuestionIndex(controller.currentQuestionIndex + 1)
	w.WriteHeader(http.StatusOK)
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	validator_host := os.Getenv("VALIDATOR_HOST")
	epvalidator := fmt.Sprintf("http://%s:8096/check_config", validator_host)
	resp, err := http.Post(epvalidator, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(respBody))

	// Check if the output contains the word "success"
	if resp.StatusCode == http.StatusOK {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Accepted"))

		index, err := strconv.Atoi(question.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		controller.questionTrialsLeft[controller.currentQuestionIndex-1] = controller.questionTrialsLeft[controller.currentQuestionIndex-1]
		controller.completedQuestionTrials[controller.currentQuestionIndex-1] = controller.questionTrialsLeft[controller.currentQuestionIndex-1]
		controller.SetCurrentQuestionIndex(index + 1)
		w.WriteHeader(http.StatusOK)
	} else {
		controller.questionTrialsLeft[controller.currentQuestionIndex-1] = controller.questionTrialsLeft[controller.currentQuestionIndex-1] - 1
		errormsg := fmt.Sprintf("%v", string(respBody))
		http.Error(w, errormsg, http.StatusBadRequest)
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

	reqBody, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	gitkatas_host := os.Getenv("GITKATAS_HOST")
	epgitkatas := fmt.Sprintf("http://%s:8095/stage_before_actions", gitkatas_host)
	resp, err := http.Post(epgitkatas, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	questionDetails := controller.currentCuestionDetails(question.TypeQuestion, question)

	if resp.StatusCode == http.StatusOK {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(questionDetails)
	} else {
		http.Error(w, "Output does not contain 'accepted'", http.StatusBadRequest)
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
