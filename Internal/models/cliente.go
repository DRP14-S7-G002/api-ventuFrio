package models

type Cliente struct {
	Id        int    `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	CPF       string `json:"cpf"`
	Telefone  string `json:"contato_celular"`
	Email     string `json:"email"`
	Endereco  string `json:"endereco"`
}
