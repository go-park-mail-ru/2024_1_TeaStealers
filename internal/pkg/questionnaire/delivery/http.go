package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/questionnaire"
	"2024_1_TeaStealers/internal/pkg/utils"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
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
	data := vars["theme"]
	th := models.QuestionTheme(data)
	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	questions, err := h.uc.GetQuestionsByTheme(&th)
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

	statistics, err := h.uc.GetAnswerStatistics()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, statistics); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
