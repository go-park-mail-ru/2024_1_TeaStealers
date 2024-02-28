package delivery

import (
	"2024_1_TeaStealers/internal/pkg/auth"
	"net/http"
)

type Handler struct {
	usecase auth.Usecase
}

func NewHandler(usecase auth.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

}
