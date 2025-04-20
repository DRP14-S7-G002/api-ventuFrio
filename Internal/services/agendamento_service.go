package services

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/repositories"
)

type AgendamentoService interface {
	GetAll() ([]models.Agendamento, error)
	GetByID(id int) (models.Agendamento, error)
	Create(agendamento models.Agendamento) (models.Agendamento, error)
	Update(agendamento models.Agendamento) (models.Agendamento, error)
	Delete(agendamento models.Agendamento) error
}

type agendamentoService struct {
	repo repositories.AgendamentoRepository
}

func NewAgendamentoService(repo repositories.AgendamentoRepository) AgendamentoService {
	return &agendamentoService{repo}
}

func (s *agendamentoService) GetAll() ([]models.Agendamento, error) {
	return s.repo.GetAll()
}

func (s *agendamentoService) GetByID(id int) (models.Agendamento, error) {
	return s.repo.GetByID(id)
}

func (s *agendamentoService) Create(agendamento models.Agendamento) (models.Agendamento, error) {
	return s.repo.Create(agendamento)
}

func (s *agendamentoService) Update(agendamento models.Agendamento) (models.Agendamento, error) {
	return s.repo.Update(agendamento)
}

func (s *agendamentoService) Delete(agendamento models.Agendamento) error {
	return s.repo.Delete(agendamento)
}
