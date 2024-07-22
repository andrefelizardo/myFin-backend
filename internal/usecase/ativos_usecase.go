package usecase

import (
	"github.com/andrefelizard/myFin-backend/internal/model"
	"github.com/andrefelizard/myFin-backend/internal/repository"
)

type AtivosUseCase struct {
	repository repository.AtivosRepository
}

func NewAtivosUseCase(repo repository.AtivosRepository) AtivosUseCase {
	return AtivosUseCase{
		repository: repo,
	}
}

func (a *AtivosUseCase) GetAtivos() ([]model.Ativo, error) {
	return a.repository.GetAtivos()
}