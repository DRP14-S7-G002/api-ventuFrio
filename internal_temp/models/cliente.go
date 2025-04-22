package models

import "time"

type Cliente struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Nome      string    `gorm:"size:150;not null" json:"nome"`
	Sobrenome string    `gorm:"size:150;not null" json:"sobrenome"`
	Telefone  string    `gorm:"size:15" json:"telefone"`
	CPF       string    `gorm:"size:11;unique" json:"cpf"`
	Rua       string    `gorm:"size:150" json:"rua"`
	Numero    string    `gorm:"size:45" json:"numero"`
	Bairro    string    `gorm:"size:100" json:"bairro"`
	CEP       string    `gorm:"size:8" json:"cep"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	LogDelete string    `gorm:"size:300" json:"log_delete"`

	Orcamentos []Orcamento `gorm:"foreignKey:ClienteID" json:"orcamentos"`
}
