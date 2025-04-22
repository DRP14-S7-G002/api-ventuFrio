package repository

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"gorm.io/gorm"
)

type OrdemDeServicoRepository interface {
	GetAll() ([]models.OrdemDeServico, error)
	GetByID(id int) (models.OrdemDeServico, error)
	Create(ordem models.OrdemDeServico) error
	Update(id int, ordem models.OrdemDeServico) error
	Delete(id int) error
}

type ordemDeServicoRepo struct {
	db *gorm.DB
}

func NewOrdemDeServicoRepository(db *gorm.DB) OrdemDeServicoRepository {
	return &ordemDeServicoRepo{db}
}

func (r *ordemDeServicoRepo) GetAll() ([]models.OrdemDeServico, error) {
	var ordens []models.OrdemDeServico
	err := r.db.Preload("Budget").Find(&ordens).Error
	return ordens, err
}

func (r *ordemDeServicoRepo) GetByID(id int) (models.OrdemDeServico, error) {
	var ordem models.OrdemDeServico
	err := r.db.First(&ordem, id).Error
	return ordem, err
}

func (r *ordemDeServicoRepo) Create(ordem models.OrdemDeServico) error {
	return r.db.Create(&ordem).Error
}

func (r *ordemDeServicoRepo) Update(id int, updated models.OrdemDeServico) error {
	return r.db.Model(&models.OrdemDeServico{}).Where("id = ?", id).Updates(updated).Error
}

func (r *ordemDeServicoRepo) Delete(id int) error {
	return r.db.Delete(&models.OrdemDeServico{}, id).Error
}
