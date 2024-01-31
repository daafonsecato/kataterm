package models

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"gitkatas/pkg/db"
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

func (store *QuestionStore) GetQuestion(id string) (*Question, error) {
	var question Question
	err := store.db.QueryRow("SELECT ID, Content_Text, Hint, Subtext, Type_Question, Staging_Message, Options, Before_Actions, Answer, Test_spec_filename, Trials FROM questions WHERE ID = $1", id).Scan(&question.ID, &question.ContentText, &question.Hint, &question.Subtext, &question.TypeQuestion, &question.StagingMessage, &question.Options, &question.BeforeActions, &question.Answer, &question.TestSpecFilename, &question.Trials)
	if err != nil {
		return nil, fmt.Errorf("could not get question: %v", err)
	}
	return &question, nil
}
