package repositories

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"

	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/models"
)

type OrcamentoRepository interface {
	GetAll() ([]models.Orcamento, error)
	GetByID(id int) (models.Orcamento, error)
	Create(orcamento models.Orcamento) (models.Orcamento, error)
	Update(orcamento models.Orcamento) (models.Orcamento, error)
	Delete(orcamento models.Orcamento) error
}

type orcamentoRepository struct{}

func NewOrcamentoRepository() OrcamentoRepository {
	return &orcamentoRepository{}
}

func (r *orcamentoRepository) GetAll() ([]models.Orcamento, error) {
	var orcamentos []models.Orcamento
	err := db.DB.Find(&orcamentos).Error
	return orcamentos, err
}

func (r *orcamentoRepository) GetByID(id int) (models.Orcamento, error) {
	var orcamento models.Orcamento
	err := db.DB.First(&orcamento, id).Error
	return orcamento, err
}

func (r *orcamentoRepository) Create(orcamento models.Orcamento) (models.Orcamento, error) {
	err := db.DB.Create(&orcamento).Error
	return orcamento, err
}

func (r *orcamentoRepository) Update(orcamento models.Orcamento) (models.Orcamento, error) {
	err := db.DB.Save(&orcamento).Error
	return orcamento, err
}

func (r *orcamentoRepository) Delete(orcamento models.Orcamento) error {
	err := db.DB.Delete(&orcamento).Error
	return err
}
