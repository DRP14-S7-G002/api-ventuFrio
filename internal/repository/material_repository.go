package repository

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"gorm.io/gorm"
)

type MaterialRepository interface {
	GetAll() ([]models.Material, error)
	GetByID(id int) (models.Material, error)
	Create(material models.Material) error
	Update(id int, material models.Material) error
	Delete(id int) error
}

type materialRepo struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepo{db}
}

func (r *materialRepo) GetAll() ([]models.Material, error) {
	var materiais []models.Material
	err := r.db.Preload("OrdemDeServico").Find(&materiais).Error
	return materiais, err
}

func (r *materialRepo) GetByID(id int) (models.Material, error) {
	var material models.Material
	err := r.db.First(&material, id).Error
	return material, err
}

func (r *materialRepo) Create(material models.Material) error {
	return r.db.Create(&material).Error
}

func (r *materialRepo) Update(id int, updated models.Material) error {
	return r.db.Model(&models.Material{}).Where("id = ?", id).Updates(updated).Error
}

func (r *materialRepo) Delete(id int) error {
	return r.db.Delete(&models.Material{}, id).Error
}
