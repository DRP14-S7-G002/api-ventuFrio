package models

import "time"

type Agendamento struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	DataVisita time.Time `json:"data_visita"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Orcamentos []Orcamento `gorm:"foreignKey:AgendamentoID" json:"orcamentos"`
}
