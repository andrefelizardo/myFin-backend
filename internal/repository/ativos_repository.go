package repository

import (
	"database/sql"
	"fmt"

	"github.com/andrefelizard/myFin-backend/internal/model"
)

type AtivosRepository struct {
	connection *sql.DB
}

func NewAtivosRepository(connection *sql.DB) AtivosRepository {
	return AtivosRepository{
		connection: connection,
	}
}

func (ar *AtivosRepository) GetAtivos() ([]model.Ativo, error) {
	query := "SELECT * FROM ativos"
	rows, err := ar.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Ativo{}, err
	}

	var ativos []model.Ativo
	var ativo model.Ativo

	for rows.Next() {
		err = rows.Scan(&ativo.Id, &ativo.Nome, &ativo.Quantidade)
		if err != nil {
			fmt.Println(err)
			return []model.Ativo{}, err
		}

		ativos = append(ativos, ativo)
	}

	rows.Close()

	return ativos, nil
}