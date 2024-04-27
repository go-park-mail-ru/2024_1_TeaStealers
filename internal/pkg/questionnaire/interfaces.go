//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}

package questionnaire

import (
	"2024_1_TeaStealers/internal/models"
	"context"
)

// QuestionnaireUsecase represents the usecase interface for questions iframes.
type QuestionnaireUsecase interface {
	GetQuestionsByTheme(context.Context, *models.QuestionTheme) ([]*models.QuestionResp, error)
	UploadAnswer(context.Context, *models.QuestionAnswerResp) error
	GetAnswerStatistics(ctx context.Context) ([]*models.ThemeStatistic, error)
}

// QuestionnaireRepo represents the repository interface for questions iframes.
type QuestionnaireRepo interface {
	SelectQuestionsByTheme(context.Context, *models.QuestionTheme) ([]*models.QuestionResp, error)
	StoreAnswer(context.Context, *models.QuestionAnswer) error
	SelectAnswerStatistics(context.Context) ([]*models.ThemeStatistic, error)
}
