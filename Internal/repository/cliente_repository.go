package repository

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"gorm.io/gorm"
)

type ClienteRepository interface {
	GetAll() ([]models.Cliente, error)
	GetByID(id int) (models.Cliente, error)
	Create(cliente models.Cliente) error
	Update(id int, cliente models.Cliente) error
	Delete(id int) error
}

type clienteRepo struct {
	db *gorm.DB
}

func NewClienteRepository(db *gorm.DB) ClienteRepository {
	return &clienteRepo{db}
}

func (r *clienteRepo) GetAll() ([]models.Cliente, error) {
	var clientes []models.Cliente
	err := r.db.Find(&clientes).Error
	return clientes, err
}

func (r *clienteRepo) GetByID(id int) (models.Cliente, error) {
	var cliente models.Cliente
	err := r.db.First(&cliente, id).Error
	return cliente, err
}

func (r *clienteRepo) Create(cliente models.Cliente) error {
	return r.db.Create(&cliente).Error
}

func (r *clienteRepo) Update(id int, updated models.Cliente) error {
	return r.db.Model(&models.Cliente{}).Where("id = ?", id).Updates(updated).Error
}

func (r *clienteRepo) Delete(id int) error {
	return r.db.Delete(&models.Cliente{}, id).Error
}
