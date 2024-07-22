package model

type Ativo struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Quantidade float64 `json:"quantidade"`
}