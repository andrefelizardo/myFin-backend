package usecase

import (
	"github.com/andrefelizard/myFin-backend/internal/model"
	"github.com/andrefelizard/myFin-backend/internal/repository"
)

type MinhasMetasUseCase struct {
	repository repository.MinhasMetasRepository
}

func NewMinhasMetasUseCase(repo repository.MinhasMetasRepository) MinhasMetasUseCase {
	return MinhasMetasUseCase{
		repository: repo,
	}
}

func (mu *MinhasMetasUseCase) GetMinhasMetas(id string) (model.MinhasMetas, error) {
	return mu.repository.GetMinhasMetas(id)
}