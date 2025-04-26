package dto

type OrdemServicoCreateRequest struct {
	DescricaoServico string `json:"descricao_servico" binding:"required"`
	Status           string `json:"status"`
	Responsavel      string `json:"responsavel"`
	OrcamentoID      int    `json:"orcamento_id" binding:"required"`
}

type OrdemServicoResponse struct {
	ID               int    `json:"id"`
	DescricaoServico string `json:"descricao_servico"`
	Status           string `json:"status"`
	Responsavel      string `json:"responsavel"`
	OrcamentoID      int    `json:"orcamento_id"`
}
