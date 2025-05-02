package models

import "time"

type OrdemDeServico struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	DescricaoServico string    `gorm:"size:300" json:"descricao_servico"`
	Status           string    `gorm:"size:45" json:"status"`
	Responsavel      string    `gorm:"size:45" json:"responsavel"`
	OrcamentoID      int       `json:"orcamento_id"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Orcamento Orcamento  `gorm:"foreignKey:OrcamentoID" json:"orcamento"`
	Materiais []Material `gorm:"foreignKey:OrdemDeServicoID" json:"materiais"`
}

func (OrdemDeServico) TableName() string {
	return "ordem_servico"
}
