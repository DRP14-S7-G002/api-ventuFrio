package services

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/repositories"
)

type EnderecoService interface {
	GetAll() ([]models.Endereco, error)
	GetByID(id int) (models.Endereco, error)
	Create(endereco models.Endereco) (models.Endereco, error)
	Update(endereco models.Endereco) (models.Endereco, error)
	Delete(endereco models.Endereco) error
}

type enderecoService struct {
	repo repositories.EnderecoRepository
}

func NewEnderecoService(repo repositories.EnderecoRepository) EnderecoService {
	return &enderecoService{repo}
}

func (s *enderecoService) GetAll() ([]models.Endereco, error) {
	return s.repo.GetAll()
}

func (s *enderecoService) GetByID(id int) (models.Endereco, error) {
	return s.repo.GetByID(id)
}

func (s *enderecoService) Create(endereco models.Endereco) (models.Endereco, error) {
	return s.repo.Create(endereco)
}

func (s *enderecoService) Update(endereco models.Endereco) (models.Endereco, error) {
	return s.repo.Update(endereco)
}

func (s *enderecoService) Delete(endereco models.Endereco) error {
	return s.repo.Delete(endereco)
}
