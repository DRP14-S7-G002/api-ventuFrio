package repository

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"gorm.io/gorm"
)

type AgendamentoRepository interface {
	GetAll() ([]models.Agendamento, error)
	GetByID(id int) (models.Agendamento, error)
	Create(agendamento models.Agendamento) error
	Update(id int, agendamento models.Agendamento) error
	Delete(id int) error
}

type agendamentoRepo struct {
	db *gorm.DB
}

func NewAgendamentoRepository(db *gorm.DB) AgendamentoRepository {
	return &agendamentoRepo{db}
}

func (r *agendamentoRepo) GetAll() ([]models.Agendamento, error) {
	var agendamentos []models.Agendamento
	err := r.db.Find(&agendamentos).Error
	return agendamentos, err
}

func (r *agendamentoRepo) GetByID(id int) (models.Agendamento, error) {
	var agendamento models.Agendamento
	err := r.db.First(&agendamento, id).Error
	return agendamento, err
}

func (r *agendamentoRepo) Create(agendamento models.Agendamento) error {
	return r.db.Create(&agendamento).Error
}

func (r *agendamentoRepo) Update(id int, updated models.Agendamento) error {
	return r.db.Model(&models.Agendamento{}).Where("id = ?", id).Updates(updated).Error
}

func (r *agendamentoRepo) Delete(id int) error {
	return r.db.Delete(&models.Agendamento{}, id).Error
}
