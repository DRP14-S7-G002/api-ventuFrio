package service

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/repository"
)

type MaterialService interface {
	GetAllMateriais() ([]models.Material, error)
	GetMaterialByID(id int) (models.Material, error)
	CreateMaterial(material models.Material) error
	UpdateMaterial(id int, material models.Material) error
	DeleteMaterial(id int) error
}

type materialService struct {
	repo repository.MaterialRepository
}

func NewMaterialService(r repository.MaterialRepository) MaterialService {
	return &materialService{r}
}

func (s *materialService) GetAllMateriais() ([]models.Material, error) {
	return s.repo.GetAll()
}

func (s *materialService) GetMaterialByID(id int) (models.Material, error) {
	return s.repo.GetByID(id)
}

func (s *materialService) CreateMaterial(material models.Material) error {
	return s.repo.Create(material)
}

func (s *materialService) UpdateMaterial(id int, material models.Material) error {
	return s.repo.Update(id, material)
}

func (s *materialService) DeleteMaterial(id int) error {
	return s.repo.Delete(id)
}
