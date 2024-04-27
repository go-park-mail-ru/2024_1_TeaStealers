package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/questionnaire"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
	"go.uber.org/zap"
)

// QuestionnaireHandler handles HTTP requests for questionnaire.
type QuestionnaireHandler struct {
	// uc represents the usecase interface for questionnaire.
	uc     questionnaire.QuestionnaireUsecase
	logger *zap.Logger
}

// NewQuestionnaireHandler creates a new instance of QuestionnaireHandler.
func NewQuestionnaireHandler(uc questionnaire.QuestionnaireUsecase, logger *zap.Logger) *QuestionnaireHandler {
	return &QuestionnaireHandler{uc: uc, logger: logger}
}

// GetQuestionsByTheme handles the request for getting questions by theme
func (h *QuestionnaireHandler) GetQuestionsByTheme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	theme := models.QuestionTheme(vars["theme"])

	questions, err := h.uc.GetQuestionsByTheme(r.Context(), &theme)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, questions); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// UploadAnswer handles the request for uploading answer for question
func (h *QuestionnaireHandler) UploadAnswer(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CookieName)
	_, ok := id.(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	data := models.QuestionAnswerResp{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	err := h.uc.UploadAnswer(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, "ok"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetAnswerStatistics handles the request for getting questionnaire statistics
func (h *QuestionnaireHandler) GetAnswerStatistics(w http.ResponseWriter, r *http.Request) {

	statistics, err := h.uc.GetAnswerStatistics(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, statistics); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
