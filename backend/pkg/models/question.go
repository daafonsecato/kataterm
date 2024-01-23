package models

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/david8128/quizard-backend/pkg/db"
)

type Question struct {
	ID               string          `json:"id"`
	ContentText      string          `json:"content_text"`
	Hint             string          `json:"hint"`
	Subtext          string          `json:"subtext"`
	TypeQuestion     string          `json:"type_question"`
	StagingMessage   string          `json:"staging_message"`
	Options          json.RawMessage `json:"options"`
	BeforeActions    json.RawMessage `json:"before_actions"`
	Answer           string          `json:"answer"`
	TestSpecFilename string          `json:"test_spec_filename"`
	Trials           string          `json:"trials"`
}

type QuestionStore struct {
	db *sql.DB
}

func NewQuestionStore(database *sql.DB) *QuestionStore {
	db.InitDB()
	db, err := db.InitDB()
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to open database: %v", err)
		panic(errorMsg)
	}

	err = db.Ping()
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to ping database: %v", err)
		panic(errorMsg)
	}

	if db == nil {
		panic("db is nil")
	}
	return &QuestionStore{
		db: db,
	}
}

func (store *QuestionStore) Close() {
	store.db.Close()
}

func (store *QuestionStore) GetNumberOfQuestions() (int, error) {
	var count int
	err := store.db.QueryRow("SELECT COUNT(*) FROM questions").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not get number of questions: %v", err)
	}
	return count, nil
}

func (store *QuestionStore) GetQuestions() ([]*Question, error) {
	rows, err := store.db.Query("SELECT ID, Content_Text, Hint, Subtext, Type_Question, Staging_Message, Options, Before_Actions, Answer, Test_spec_filename, Trials FROM questions")
	if err != nil {
		return nil, fmt.Errorf("could not get questions: %v", err)
	}
	defer rows.Close()

	var questions []*Question
	for rows.Next() {
		var question Question
		err := rows.Scan(&question.ID, &question.ContentText, &question.Hint, &question.Subtext, &question.TypeQuestion, &question.StagingMessage, &question.Options, &question.BeforeActions, &question.Answer, &question.TestSpecFilename, &question.Trials)
		if err != nil {
			return nil, fmt.Errorf("could not scan question: %v", err)
		}
		questions = append(questions, &question)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate over questions: %v", err)
	}

	return questions, nil
}

func (store *QuestionStore) GetQuestion(id string) (*Question, error) {
	var question Question
	err := store.db.QueryRow("SELECT ID, Content_Text, Hint, Subtext, Type_Question, Staging_Message, Options, Before_Actions, Answer, Test_spec_filename, Trials FROM questions WHERE ID = $1", id).Scan(&question.ID, &question.ContentText, &question.Hint, &question.Subtext, &question.TypeQuestion, &question.StagingMessage, &question.Options, &question.BeforeActions, &question.Answer, &question.TestSpecFilename, &question.Trials)
	if err != nil {
		return nil, fmt.Errorf("could not get question: %v", err)
	}
	return &question, nil
}

func (store *QuestionStore) CreateQuestion(id string, contentText string, hint string, subtext string, typeQuestion string, stagingMessage string, options json.RawMessage, beforeActions json.RawMessage, answer string, testSpecFilename string, trials string) (*Question, error) {
	_, err := store.db.Exec("INSERT INTO questions(ID, Content_Text, Hint, Subtext, Type_Question, Staging_Message, Options, Before_Actions, Answer, Test_spec_filename, Trials) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", id, contentText, hint, subtext, typeQuestion, stagingMessage, options, beforeActions, answer, testSpecFilename, trials)
	if err != nil {
		return nil, fmt.Errorf("could not create question: %v", err)
	}
	return &Question{ID: id, ContentText: contentText, Hint: hint, Subtext: subtext, TypeQuestion: typeQuestion, StagingMessage: stagingMessage, Options: options, BeforeActions: beforeActions, Answer: answer, TestSpecFilename: testSpecFilename, Trials: trials}, nil
}

func (store *QuestionStore) UpdateQuestion(id string, contentText string, hint string, subtext string, typeQuestion string, stagingMessage string, options json.RawMessage, beforeActions json.RawMessage, answer string, testSpecFilename string, trials string) error {
	_, err := store.db.Exec("UPDATE questions SET Content_Text = $1, Hint = $2, Subtext = $3, Type_Question = $4, Staging_Message = $5, Options = $6, Before_Actions = $7, Answer = $8, Test_spec_filename = $9, Trials = $10 WHERE ID = $11", contentText, hint, subtext, typeQuestion, stagingMessage, options, beforeActions, answer, testSpecFilename, trials, id)
	if err != nil {
		return fmt.Errorf("could not update question: %v", err)
	}
	return nil
}

func (store *QuestionStore) DeleteQuestion(id string) error {
	_, err := store.db.Exec("DELETE FROM questions WHERE ID = $1", id)
	if err != nil {
		return fmt.Errorf("could not delete question: %v", err)
	}
	return nil
}
