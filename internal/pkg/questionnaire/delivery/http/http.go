package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/questionnaire"
	question "2024_1_TeaStealers/internal/pkg/questionnaire"
	genQuestion "2024_1_TeaStealers/internal/pkg/questionnaire/delivery/grpc/gen"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// QuestionnaireClientHandler handles HTTP requests for questions.
type QuestionnaireClientHandler struct {
	client genQuestion.QuestionClient
	uc     question.QuestionnaireUsecase
	logger *zap.Logger
}

// NewQuestionnaireClientHandler creates a new instance of QuestionnaireClientHandler.
func NewQuestionnaireClientHandler(grpcConn *grpc.ClientConn, uc questionnaire.QuestionnaireUsecase, logger *zap.Logger) *QuestionnaireClientHandler {
	return &QuestionnaireClientHandler{client: genQuestion.NewQuestionClient(grpcConn), uc: uc, logger: logger}
}

// GetQuestionsByTheme handles the request for getting questions by theme
func (h *QuestionnaireClientHandler) GetQuestionsByTheme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	theme := models.QuestionTheme(vars["theme"])

	questions, err := h.client.GetQuestionsByTheme(r.Context(), &genQuestion.GetQuestionsByThemeRequest{Theme: string(theme)})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, questions); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// UploadAnswer handles the request for uploading answer for question
func (h *QuestionnaireClientHandler) UploadAnswer(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CookieName)
	ID, ok := id.(int64)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	data := models.QuestionAnswerResp{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	_, err := h.client.UploadAnswer(r.Context(), &genQuestion.UploadAnswerRequest{UserId: ID, QuestionId: data.QuestionID, Mark: int64(data.Mark)})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, "ok"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetAnswerStatistics handles the request for getting questionnaire statistics
func (h *QuestionnaireClientHandler) GetAnswerStatistics(w http.ResponseWriter, r *http.Request) {

	statistics, err := h.uc.GetAnswerStatistics(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, statistics); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
