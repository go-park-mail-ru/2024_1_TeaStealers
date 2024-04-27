package models

import (
	"time"

	"github.com/satori/uuid"
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
	Title string `json:"title"`
	// CountAnswers is count of answers users with such answer for the question.
	CountAnswers int `json:"count_answers"`
	// Mark is quantity of users with such answer for the question.
	Mark int `json:"mark"`
}

type ThemeStatistic struct {
	// Theme is the theme for the questions.
	Theme QuestionTheme `json:"theme"`
	// Questions contains filenames of photos for advert.
	Questions []*QuestionAnswerStatisticsResp `json:"questions"`
}
