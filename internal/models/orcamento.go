package models

import "time"

type Orcamento struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	DescricaoInicial string    `gorm:"size:300" json:"descricao_inicial"`
	DescricaoItem    string    `gorm:"size:300" json:"descricao_item"`
	Status           string    `gorm:"size:45" json:"status"`
	PrazoEntrega     time.Time `json:"prazo_entrega"`
	Valor            float64   `gorm:"type:decimal(10,2)" json:"valor"`
	ClienteID        int       `json:"cliente_id"`
	AgendamentoID    int       `json:"agendamento_id"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeleteLog        string    `gorm:"size:300" json:"delete_log"`

	Cliente         Cliente          `gorm:"foreignKey:ClienteID" json:"cliente"`
	Agendamento     Agendamento      `gorm:"foreignKey:AgendamentoID" json:"agendamento"`
	OrdensDeServico []OrdemDeServico `gorm:"foreignKey:OrcamentoID" json:"ordens_de_servico"`
}

func (Orcamento) TableName() string {
	return "orcamento"
}
