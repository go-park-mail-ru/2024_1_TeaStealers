package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/questionnaire"
	"context"

	"go.uber.org/zap"
)

// QuestionnaireUsecase represents the usecase.
type QuestionnaireUsecase struct {
	repo   questionnaire.QuestionnaireRepo
	logger *zap.Logger
}

// NewQuestionnaireUsecase creates a new instance of QuestionnaireUsecase.
func NewQuestionnaireUsecase(repo questionnaire.QuestionnaireRepo, logger *zap.Logger) *QuestionnaireUsecase {
	return &QuestionnaireUsecase{repo: repo, logger: logger}
}

// GetQuestionsByTheme handles the creation advert process.
func (u *QuestionnaireUsecase) GetQuestionsByTheme(ctx context.Context, theme *models.QuestionTheme) ([]*models.QuestionResp, error) {
	qResp, err := u.repo.SelectQuestionsByTheme(ctx, theme)
	if err != nil {
		return nil, err
	}

	return qResp, nil
}

// UploadAnswer handles the creation of answer.
func (u *QuestionnaireUsecase) UploadAnswer(ctx context.Context, answ *models.QuestionAnswerResp, userId int64) error {
	err := u.repo.StoreAnswer(ctx, &models.QuestionAnswer{QuestionID: answ.QuestionID, UserID: userId, Mark: answ.Mark})
	return err
}

// GetAnswerStatistics handles the creation of answer.
func (u *QuestionnaireUsecase) GetAnswerStatistics(ctx context.Context) ([]*models.ThemeStatistic, error) {
	statistic, err := u.repo.SelectAnswerStatistics(ctx)
	return statistic, err
}
