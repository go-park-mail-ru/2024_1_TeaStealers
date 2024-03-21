package usecase

import (
	"2024_1_TeaStealers/internal/pkg/complexes"
)

// ComplexUsecase represents the usecase for complex using.
type ComplexUsecase struct {
	repo complexes.ComplexRepo
}

// NewComplexUsecase creates a new instance of ComplexUsecase.
func NewComplexUsecase(repo complexes.ComplexRepo) *ComplexUsecase {
	return &ComplexUsecase{repo: repo}
}
