package repositories

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"

	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/models"
)

type EnderecoRepository interface {
	GetAll() ([]models.Endereco, error)
	GetByID(id int) (models.Endereco, error)
	Create(endereco models.Endereco) (models.Endereco, error)
	Update(endereco models.Endereco) (models.Endereco, error)
	Delete(endereco models.Endereco) error
}

type enderecoRepository struct{}

func NewEnderecoRepository() EnderecoRepository {
	return &enderecoRepository{}
}

func (r *enderecoRepository) GetAll() ([]models.Endereco, error) {
	var enderecos []models.Endereco
	err := db.DB.Find(&enderecos).Error
	return enderecos, err
}

func (r *enderecoRepository) GetByID(id int) (models.Endereco, error) {
	var endereco models.Endereco
	err := db.DB.First(&endereco, id).Error
	return endereco, err
}

func (r *enderecoRepository) Create(endereco models.Endereco) (models.Endereco, error) {
	err := db.DB.Create(&endereco).Error
	return endereco, err
}

func (r *enderecoRepository) Update(endereco models.Endereco) (models.Endereco, error) {
	err := db.DB.Save(&endereco).Error
	return endereco, err
}

func (r *enderecoRepository) Delete(endereco models.Endereco) error {
	err := db.DB.Delete(&endereco).Error
	return err
}
