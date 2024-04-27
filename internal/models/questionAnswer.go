package models

import (
	"github.com/satori/uuid"
	"time"
)

// QuestionAnswer represents a user answer for question.
type QuestionAnswer struct {
	// UserID is the unique identifier for the user.
	UserID uuid.UUID `json:"user_id"`
	// QuestionID is the unique identifier for the question.
	QuestionID uuid.UUID `json:"question_id"`
	// Mark is user mark for the question.
	Mark int `json:"mark"`
	// DateCreation is the date when the price change was created.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the price change is deleted.
	IsDeleted bool `json:"-"`
}

type QuestionAnswerResp struct {
	// UserID is the unique identifier for the user.
	UserID uuid.UUID `json:"user_id"`
	// QuestionID is the unique identifier for the question.
	QuestionID uuid.UUID `json:"question_id"`
	// Mark is user mark for the question.
	Mark int `json:"mark"`
}

type QuestionAnswerStatisticsResp struct {
	// QuestionID is the unique identifier for the question.
	QuestionID uuid.UUID `json:"question_id"`
	// NumberMark is quantity of users with such answer for the question.
	NumberMark int `json:"mark"`
}
