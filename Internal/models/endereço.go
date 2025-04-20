package models

type Endereco struct {
	Id         int    `json:"id"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Numero     string `json:"numero"`
	Cidade     string `json:"cidade"`
	Estado     string `json:"estado"`
	CEP        string `json:"cep"`
}
