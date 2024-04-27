//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}

package questionnaire

import (
	"2024_1_TeaStealers/internal/models"
)

// QuestionnaireUsecase represents the usecase interface for questions iframes.
type QuestionnaireUsecase interface {
	GetQuestionsByTheme(*models.QuestionTheme) ([]*models.QuestionResp, error)
	UploadAnswer(*models.QuestionAnswerResp) error
	GetAnswerStatistics() ([]*models.QuestionAnswerResp, error)
}

// QuestionnaireRepo represents the repository interface for questions iframes.
type QuestionnaireRepo interface {
	SelectQuestionsByTheme(*models.QuestionTheme) ([]*models.QuestionResp, error)
	StoreAnswer(*models.QuestionAnswerResp) error
	SelectAnswerStatistics() ([]*models.QuestionAnswerResp, error)
}
