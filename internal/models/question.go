package models

import (
	"github.com/satori/uuid"
	"time"
)

type QuestionTheme string

const (
	// mainPageTheme represents questions for mainPage page.
	MainPageTheme QuestionTheme = "mainPage"
	// createAdvertTheme represents questions for createAdvert page.
	CreateAdvertTheme QuestionTheme = "createAdvert"
	// filterPageTheme represents questions for filterPage page.
	FilterPageTheme QuestionTheme = "filterPage"
	// profileTheme represents questions for profile page.
	ProfileTheme QuestionTheme = "profile"
	// myAdvertsTheme represents questions for myAdverts page.
	MyAdvertsTheme QuestionTheme = "myAdverts"
)

// Question represents a question for user.
type Question struct {
	// ID is the unique identifier for the question.
	ID uuid.UUID `json:"id"`
	// QuestionText is the text of question which user will be asked.
	QuestionText string `json:"question_text"`
	// QuestionTheme represents page where question will be asked.
	Theme QuestionTheme `json:"question_theme"`
	// MaxMark is the max possible mark of the question.
	MaxMark int64 `json:"max_mark"`
	// DateCreation is the date when the price change was created.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the price change is deleted.
	IsDeleted bool `json:"-"`
}

type QuestionResp struct {
	// ID is the unique identifier for the question.
	ID uuid.UUID `json:"id"`
	// QuestionText is the text of question which user will be asked.
	QuestionText string `json:"question_text"`
	// QuestionTheme represents page where question will be asked.
	MaxMark int64 `json:"max_mark"`
}
