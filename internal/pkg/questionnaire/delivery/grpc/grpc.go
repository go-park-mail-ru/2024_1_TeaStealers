package grpc

import (
	"2024_1_TeaStealers/internal/models"
	question "2024_1_TeaStealers/internal/pkg/questionnaire"
	genQuestion "2024_1_TeaStealers/internal/pkg/questionnaire/delivery/grpc/gen"
	"context"
	"log"

	"go.uber.org/zap"
)

// QuestionServerHandler handles HTTP requests for questions.
type QuestionServerHandler struct {
	genQuestion.QuestionServer
	// uc represents the usecase interface for authentication.
	uc     question.QuestionnaireUsecase
	logger *zap.Logger
}

// NewQuestionServerHandler creates a new instance of QuestionServerHandler.
func NewQuestionServerHandler(uc question.QuestionnaireUsecase, logger *zap.Logger) *QuestionServerHandler {
	return &QuestionServerHandler{uc: uc, logger: logger}
}

func (h *QuestionServerHandler) GetQuestionsByTheme(ctx context.Context, req *genQuestion.GetQuestionsByThemeRequest) (*genQuestion.GetQuestionsByThemeResponse, error) {

	questions, err := h.uc.GetQuestionsByTheme(ctx, (*models.QuestionTheme)(&req.Theme))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response := &genQuestion.GetQuestionsByThemeResponse{}
	response.Questions = make([]*genQuestion.QuestionResp, 0, len(questions))
	for _, question := range questions {
		response.Questions = append(response.Questions, &genQuestion.QuestionResp{Id: question.ID, QuestionText: question.QuestionText, MaxMark: question.MaxMark})
	}

	return response, nil
}

func (h *QuestionServerHandler) UploadAnswer(ctx context.Context, req *genQuestion.UploadAnswerRequest) (*genQuestion.UploadAnswerResponse, error) {

	data := &models.QuestionAnswerResp{QuestionID: req.QuestionId, Mark: int(req.Mark)}

	err := h.uc.UploadAnswer(ctx, data, req.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &genQuestion.UploadAnswerResponse{}, nil
}
