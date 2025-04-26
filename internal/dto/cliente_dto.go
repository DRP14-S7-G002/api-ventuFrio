package dto

type ClienteCreateRequest struct {
	Nome     string `json:"nome" binding:"required"`
	Telefone string `json:"telefone"`
	CPF      string `json:"cpf" binding:"required"`
	Rua      string `json:"rua"`
	Numero   string `json:"numero"`
	Bairro   string `json:"bairro"`
	CEP      string `json:"cep"`
}

type ClienteResponse struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	CPF      string `json:"cpf"`
	Rua      string `json:"rua"`
	Numero   string `json:"numero"`
	Bairro   string `json:"bairro"`
	CEP      string `json:"cep"`
}
