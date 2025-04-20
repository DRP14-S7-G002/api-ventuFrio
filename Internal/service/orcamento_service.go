package service

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/repository"
)

type OrcamentoService interface {
	GetAllOrcamentos() ([]models.Orcamento, error)
	GetOrcamentoByID(id int) (models.Orcamento, error)
	CreateOrcamento(orcamento models.Orcamento) error
	UpdateOrcamento(id int, orcamento models.Orcamento) error
	DeleteOrcamento(id int) error
}

type orcamentoService struct {
	repo repository.OrcamentoRepository
}

func NewOrcamentoService(r repository.OrcamentoRepository) OrcamentoService {
	return &orcamentoService{r}
}

func (s *orcamentoService) GetAllOrcamentos() ([]models.Orcamento, error) {
	return s.repo.GetAll()
}

func (s *orcamentoService) GetOrcamentoByID(id int) (models.Orcamento, error) {
	return s.repo.GetByID(id)
}

func (s *orcamentoService) CreateOrcamento(orcamento models.Orcamento) error {
	return s.repo.Create(orcamento)
}

func (s *orcamentoService) UpdateOrcamento(id int, orcamento models.Orcamento) error {
	return s.repo.Update(id, orcamento)
}

func (s *orcamentoService) DeleteOrcamento(id int) error {
	return s.repo.Delete(id)
}
