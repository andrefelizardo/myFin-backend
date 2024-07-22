package model

type MinhasMetas struct {
	UserId string `json:"user_id"`
	AcoesInternacionais int `json:"acoes_internacionais"`
	AcoesNacionais int `json:"acoes_nacionais"`
	Reits int `json:"reits"`
	Fiis int `json:"fiis"`
	Criptomoedas int `json:"criptomoedas"`
	RendaFixa int `json:"renda_fixa"`
 }