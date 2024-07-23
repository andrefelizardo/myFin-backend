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

func (ar *AtivosRepository) CreateAtivo(ativo model.Ativo) (int, error) {
	query, err := ar.connection.Prepare("INSERT INTO ativos (nome, quantidade) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(ativo.Nome, ativo.Quantidade).Scan(&ativo.Id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return ativo.Id, nil
}

func (ar *AtivosRepository) GetAtivoById(id int) (*model.Ativo, error) {
	query := "SELECT * FROM ativos WHERE id = $1"
	row := ar.connection.QueryRow(query, id)

	var ativo model.Ativo
	err := row.Scan(&ativo.Id, &ativo.Nome, &ativo.Quantidade)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Ativo not found")
		}
		fmt.Println(err)
		return nil, err
	}

	return &ativo, nil
}