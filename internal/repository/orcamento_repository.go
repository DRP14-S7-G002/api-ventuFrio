package repository

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"gorm.io/gorm"
)

type OrcamentoRepository interface {
	GetAll() ([]models.Orcamento, error)
	GetByID(id int) (models.Orcamento, error)
	Create(orcamento models.Orcamento) error
	Update(id int, orcamento models.Orcamento) error
	Delete(id int) error
}

type orcamentoRepo struct {
	db *gorm.DB
}

func NewOrcamentoRepository(db *gorm.DB) OrcamentoRepository {
	return &orcamentoRepo{db}
}

func (r *orcamentoRepo) GetAll() ([]models.Orcamento, error) {
	var orcamentos []models.Orcamento
	err := r.db.Preload("Cliente").Preload("Agendamento").Find(&orcamentos).Error
	return orcamentos, err
}

func (r *orcamentoRepo) GetByID(id int) (models.Orcamento, error) {
	var orcamento models.Orcamento
	err := r.db.First(&orcamento, id).Error
	return orcamento, err
}

func (r *orcamentoRepo) Create(orcamento models.Orcamento) error {
	return r.db.Create(&orcamento).Error
}

func (r *orcamentoRepo) Update(id int, updated models.Orcamento) error {
	return r.db.Model(&models.Orcamento{}).Where("id = ?", id).Updates(updated).Error
}

func (r *orcamentoRepo) Delete(id int) error {
	return r.db.Delete(&models.Orcamento{}, id).Error
}
