package dto

type AgendamentoCreateRequest struct {
	DataVisita string `json:"data_visita" binding:"required"`
}

type AgendamentoResponse struct {
	ID         int    `json:"id"`
	DataVisita string `json:"data_visita"`
}
