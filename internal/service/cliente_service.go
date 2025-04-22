package service

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/repository"
)

type ClienteService interface {
	GetAllClientes() ([]models.Cliente, error)
	GetClienteByID(id int) (models.Cliente, error)
	CreateCliente(cliente models.Cliente) error
	UpdateCliente(id int, cliente models.Cliente) error
	DeleteCliente(id int) error
}

type clienteService struct {
	repo repository.ClienteRepository
}

func NewClienteService(r repository.ClienteRepository) ClienteService {
	return &clienteService{r}
}

func (s *clienteService) GetAllClientes() ([]models.Cliente, error) {
	return s.repo.GetAll()
}

func (s *clienteService) GetClienteByID(id int) (models.Cliente, error) {
	return s.repo.GetByID(id)
}

func (s *clienteService) CreateCliente(cliente models.Cliente) error {
	return s.repo.Create(cliente)
}

func (s *clienteService) UpdateCliente(id int, cliente models.Cliente) error {
	return s.repo.Update(id, cliente)
}

func (s *clienteService) DeleteCliente(id int) error {
	return s.repo.Delete(id)
}
