package dto

type MaterialCreateRequest struct {
	Nome             string  `json:"nome" binding:"required"`
	Quantidade       string  `json:"quantidade"`
	Valor            float64 `json:"valor"`
	OrdemDeServicoID int     `json:"ordem_de_servico_id" binding:"required"`
}

type MaterialResponse struct {
	ID               int     `json:"id"`
	Nome             string  `json:"nome"`
	Quantidade       string  `json:"quantidade"`
	Valor            float64 `json:"valor"`
	OrdemDeServicoID int     `json:"ordem_de_servico_id"`
}
