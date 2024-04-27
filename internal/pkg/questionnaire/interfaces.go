//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}

package questionnaire

import (
	"2024_1_TeaStealers/internal/models"
	"context"
)

// QuestionnaireUsecase represents the usecase interface for questions iframes.
type QuestionnaireUsecase interface {
	GetQuestionsByTheme(*models.QuestionTheme) ([]*models.QuestionResp, error)
	UploadAnswer(*models.QuestionAnswerResp) error
	GetAnswerStatistics() ([]*models.QuestionAnswerResp, error)
}

// QuestionnaireRepo represents the repository interface for questions iframes.
type QuestionnaireRepo interface {
	SelectQuestionsByTheme(context.Context, *models.QuestionTheme) ([]*models.QuestionResp, error)
	StoreAnswer(context.Context, *models.QuestionAnswerResp) error
	SelectAnswerStatistics(context.Context) ([]*models.ThemeStatistic, error)
}
