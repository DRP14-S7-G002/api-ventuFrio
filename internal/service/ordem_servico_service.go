package service

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/repository"
)

type OrdemDeServicoService interface {
	GetAllOrdens() ([]models.OrdemDeServico, error)
	GetOrdemByID(id int) (models.OrdemDeServico, error)
	CreateOrdem(ordem models.OrdemDeServico) error
	UpdateOrdem(id int, ordem models.OrdemDeServico) error
	DeleteOrdem(id int) error
}

type ordemDeServicoService struct {
	repo repository.OrdemDeServicoRepository
}

func NewOrdemDeServicoService(r repository.OrdemDeServicoRepository) OrdemDeServicoService {
	return &ordemDeServicoService{r}
}

func (s *ordemDeServicoService) GetAllOrdens() ([]models.OrdemDeServico, error) {
	return s.repo.GetAll()
}

func (s *ordemDeServicoService) GetOrdemByID(id int) (models.OrdemDeServico, error) {
	return s.repo.GetByID(id)
}

func (s *ordemDeServicoService) CreateOrdem(ordem models.OrdemDeServico) error {
	return s.repo.Create(ordem)
}

func (s *ordemDeServicoService) UpdateOrdem(id int, ordem models.OrdemDeServico) error {
	return s.repo.Update(id, ordem)
}

func (s *ordemDeServicoService) DeleteOrdem(id int) error {
	return s.repo.Delete(id)
}
