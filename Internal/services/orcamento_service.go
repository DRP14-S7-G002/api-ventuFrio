package services

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/repositories"
)

type OrcamentoService interface {
	GetAll() ([]models.Orcamento, error)
	GetByID(id int) (models.Orcamento, error)
	Create(orcamento models.Orcamento) (models.Orcamento, error)
	Update(orcamento models.Orcamento) (models.Orcamento, error)
	Delete(orcamento models.Orcamento) error
}

type orcamentoService struct {
	repo repositories.OrcamentoRepository
}

func NewOrcamentoService(repo repositories.OrcamentoRepository) OrcamentoService {
	return &orcamentoService{repo}
}

func (s *orcamentoService) GetAll() ([]models.Orcamento, error) {
	return s.repo.GetAll()
}

func (s *orcamentoService) GetByID(id int) (models.Orcamento, error) {
	return s.repo.GetByID(id)
}

func (s *orcamentoService) Create(orcamento models.Orcamento) (models.Orcamento, error) {
	return s.repo.Create(orcamento)
}

func (s *orcamentoService) Update(orcamento models.Orcamento) (models.Orcamento, error) {
	return s.repo.Update(orcamento)
}

func (s *orcamentoService) Delete(orcamento models.Orcamento) error {
	return s.repo.Delete(orcamento)
}
