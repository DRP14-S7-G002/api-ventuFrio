package service

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/repository"
)

type AgendamentoService interface {
	GetAllAgendamentos() ([]models.Agendamento, error)
	GetAgendamentoByID(id int) (models.Agendamento, error)
	CreateAgendamento(agendamento models.Agendamento) error
	UpdateAgendamento(id int, agendamento models.Agendamento) error
	DeleteAgendamento(id int) error
}

type agendamentoService struct {
	repo repository.AgendamentoRepository
}

func NewAgendamentoService(r repository.AgendamentoRepository) AgendamentoService {
	return &agendamentoService{r}
}

func (s *agendamentoService) GetAllAgendamentos() ([]models.Agendamento, error) {
	return s.repo.GetAll()
}

func (s *agendamentoService) GetAgendamentoByID(id int) (models.Agendamento, error) {
	return s.repo.GetByID(id)
}

func (s *agendamentoService) CreateAgendamento(agendamento models.Agendamento) error {
	return s.repo.Create(agendamento)
}

func (s *agendamentoService) UpdateAgendamento(id int, agendamento models.Agendamento) error {
	return s.repo.Update(id, agendamento)
}

func (s *agendamentoService) DeleteAgendamento(id int) error {
	return s.repo.Delete(id)
}
