package models

import "time"

type Agendamento struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	DataVisita string    `json:"data_visita"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	ClienteID  int       `gorm:"not null" json:"cliente_id"`

	Orcamentos []Orcamento `gorm:"foreignKey:AgendamentoID" json:"orcamentos"`
}

func (Agendamento) TableName() string {
	return "agendamento"
}
