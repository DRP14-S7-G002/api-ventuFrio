package models

import "time"

type Material struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	Nome             string    `gorm:"size:300" json:"nome"`
	Quantidade       string    `gorm:"size:45" json:"quantidade"`
	OrdemDeServicoID int       `json:"ordem_de_servico_id"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	OrdemDeServico OrdemDeServico `gorm:"foreignKey:OrdemDeServicoID" json:"ordem_de_servico"`
}
