package repositories

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
)

type ClienteRepository interface {
	GetAll() ([]models.Cliente, error)
	GetByID(id int) (models.Cliente, error)
	Create(cliente models.Cliente) (models.Cliente, error)
	Update(cliente models.Cliente) (models.Cliente, error)
	Delete(cliente models.Cliente) error
}

type clienteRepository struct{}

func NewClienteRepository() ClienteRepository {
	return &clienteRepository{}
}

func (r *clienteRepository) GetAll() ([]models.Cliente, error) {
	var clientes []models.Cliente
	err := db.DB.Find(&clientes).Error
	return clientes, err
}

func (r *clienteRepository) GetByID(id int) (models.Cliente, error) {
	var cliente models.Cliente
	err := db.DB.First(&cliente, id).Error
	return cliente, err
}

func (r *clienteRepository) Create(cliente models.Cliente) (models.Cliente, error) {
	err := db.DB.Create(&cliente).Error
	return cliente, err
}

func (r *clienteRepository) Update(cliente models.Cliente) (models.Cliente, error) {
	err := db.DB.Save(&cliente).Error
	return cliente, err
}

func (r *clienteRepository) Delete(cliente models.Cliente) error {
	err := db.DB.Delete(&cliente).Error
	return err
}
