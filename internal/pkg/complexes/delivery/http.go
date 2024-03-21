package delivery

import (
	"2024_1_TeaStealers/internal/pkg/complexes"
)

// ComplexHandler handles HTTP requests for complex changes.
type ComplexHandler struct {
	// uc represents the usecase interface for complex changes.
	uc complexes.ComplexUsecase
}

// NewComplexHandler creates a new instance of ComplexHandler.
func NewComplexHandler(uc complexes.ComplexUsecase) *ComplexHandler {
	return &ComplexHandler{uc: uc}
}

//TODO
// Ручка на создание Комплекса по компании
// Ручка на создание Здания по комплексу
// Ручка на создание объявления по зданию (Пользователь скорее всего должен иметь статус владелец компании условно)
// Ручка на получение объявлений по комплексу
// Ручка на получение информации о комплексе
