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

func (a *AtivosUseCase) CreateAtivo(ativo model.Ativo) (model.Ativo, error) {
	ativoId, err := a.repository.CreateAtivo(ativo)
	if err != nil {
		return model.Ativo{}, err
	}
	
	ativo.Id = ativoId
	return ativo, nil
}

func (a *AtivosUseCase) GetAtivoById(id int) (*model.Ativo, error) {
	ativo, err := a.repository.GetAtivoById(id)
	if err != nil {
		return nil, err
	}
	return ativo, nil
}