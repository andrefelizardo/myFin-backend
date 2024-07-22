package repository

import (
	"database/sql"
	"fmt"

	"github.com/andrefelizard/myFin-backend/internal/model"
)

type MinhasMetasRepository struct {
	connection *sql.DB
}

func NewMinhasMetasRepository(connection *sql.DB) MinhasMetasRepository {
	return MinhasMetasRepository{
		connection: connection,
	}
}

func (mr *MinhasMetasRepository) GetMinhasMetas(userId string) (model.MinhasMetas, error) {
    query := "SELECT acoes_internacionais, acoes_nacionais, reits, fiis, criptomoedas, renda_fixa FROM minhas_metas WHERE user_id = $1"
    row := mr.connection.QueryRow(query, userId)

    var minhasMetasObj model.MinhasMetas
    err := row.Scan(&minhasMetasObj.AcoesInternacionais, &minhasMetasObj.AcoesNacionais, &minhasMetasObj.Reits, &minhasMetasObj.Fiis, &minhasMetasObj.Criptomoedas, &minhasMetasObj.RendaFixa)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No rows found for user_id:", userId)
            return model.MinhasMetas{}, nil
        }
        fmt.Println(err)
        return model.MinhasMetas{}, err
    }

    return minhasMetasObj, nil
}
