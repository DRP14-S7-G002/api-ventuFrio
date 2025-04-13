package services

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/Internal/repositories"
)

//go:generate mockgen -destination=../mocks/mock_admin_service.go -package=mocks github.com/DRP14-S7-G002/api-ventuFrio/internal/services AdminService

type AdminService interface {
	GetAll() ([]models.Admin, error)
	GetByID(id string) (models.Admin, error)
	Create(admin models.Admin) (models.Admin, error)
	Update(admin models.Admin) (models.Admin, error)
	Delete(admin models.Admin) error
}

type adminService struct {
	repo repositories.AdminRepository
}

func NewAdminService(repo repositories.AdminRepository) AdminService {
	return &adminService{repo}
}

func (s *adminService) GetAll() ([]models.Admin, error) {
	return s.repo.GetAll()
}

func (s *adminService) GetByID(id string) (models.Admin, error) {
	return s.repo.GetByID(id)
}

func (s *adminService) Create(admin models.Admin) (models.Admin, error) {
	return s.repo.Create(admin)
}

func (s *adminService) Update(admin models.Admin) (models.Admin, error) {
	return s.repo.Update(admin)
}

func (s *adminService) Delete(admin models.Admin) error {
	return s.repo.Delete(admin)
}
