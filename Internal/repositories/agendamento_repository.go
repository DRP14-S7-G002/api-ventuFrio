package repositories

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
)

type AgendamentoRepository interface {
	GetAll() ([]models.Agendamento, error)
	GetByID(id int) (models.Agendamento, error)
	Create(agendamento models.Agendamento) (models.Agendamento, error)
	Update(agendamento models.Agendamento) (models.Agendamento, error)
	Delete(agendamento models.Agendamento) error
}

type agendamentoRepository struct{}

func NewAgendamentoRepository() AgendamentoRepository {
	return &agendamentoRepository{}
}

func (r *agendamentoRepository) GetAll() ([]models.Agendamento, error) {
	var agendamentos []models.Agendamento
	err := db.DB.Find(&agendamentos).Error
	return agendamentos, err
}

func (r *agendamentoRepository) GetByID(id int) (models.Agendamento, error) {
	var agendamento models.Agendamento
	err := db.DB.First(&agendamento, id).Error
	return agendamento, err
}

func (r *agendamentoRepository) Create(agendamento models.Agendamento) (models.Agendamento, error) {
	err := db.DB.Create(&agendamento).Error
	return agendamento, err
}

func (r *agendamentoRepository) Update(agendamento models.Agendamento) (models.Agendamento, error) {
	err := db.DB.Save(&agendamento).Error
	return agendamento, err
}

func (r *agendamentoRepository) Delete(agendamento models.Agendamento) error {
	err := db.DB.Delete(&agendamento).Error
	return err
}
