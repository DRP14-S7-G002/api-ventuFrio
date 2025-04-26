package dto

type OrcamentoCreateRequest struct {
	DescricaoInicial string  `json:"descricao_inicial"`
	DescricaoItem    string  `json:"descricao_item"`
	Status           string  `json:"status"`
	PrazoEntrega     string  `json:"prazo_entrega"`
	Valor            float64 `json:"valor"`
	ClienteID        int     `json:"cliente_id" binding:"required"`
	AgendamentoID    int     `json:"agendamento_id"`
}

type OrcamentoResponse struct {
	ID               int     `json:"id"`
	DescricaoInicial string  `json:"descricao_inicial"`
	DescricaoItem    string  `json:"descricao_item"`
	Status           string  `json:"status"`
	PrazoEntrega     string  `json:"prazo_entrega"`
	Valor            float64 `json:"valor"`
	ClienteID        int     `json:"cliente_id"`
	AgendamentoID    int     `json:"agendamento_id"`
}
