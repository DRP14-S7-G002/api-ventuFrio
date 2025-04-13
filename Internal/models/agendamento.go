package models

type Agendamento struct {
	ID          int    `json:"id"`
	ClienteID   int    `json:"cliente_id"`
	EnderecoID  int    `json:"endereco_id"`
	Data        string `json:"data"`
	Hora        string `json:"hora"`
	OrcamentoID int    `json:"orcamento_id"`
	Status      string `json:"status"`
}
