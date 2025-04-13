package services

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/repositories"
)

type ClienteService interface {
	GetAll() ([]models.Cliente, error)
	GetByID(id int) (models.Cliente, error)
	Create(cliente models.Cliente) (models.Cliente, error)
	Update(cliente models.Cliente) (models.Cliente, error)
	Delete(cliente models.Cliente) error
}

type clienteService struct {
	repo repositories.ClienteRepository
}

func NewClienteService(repo repositories.ClienteRepository) ClienteService {
	return &clienteService{repo}
}

func (s *clienteService) GetAll() ([]models.Cliente, error) {
	return s.repo.GetAll()
}

func (s *clienteService) GetByID(id int) (models.Cliente, error) {
	return s.repo.GetByID(id)
}

func (s *clienteService) Create(cliente models.Cliente) (models.Cliente, error) {
	return s.repo.Create(cliente)
}

func (s *clienteService) Update(cliente models.Cliente) (models.Cliente, error) {
	return s.repo.Update(cliente)
}

func (s *clienteService) Delete(cliente models.Cliente) error {
	return s.repo.Delete(cliente)
}
