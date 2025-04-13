package repositories

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
)

type AdminRepository interface {
	GetAll() ([]models.Admin, error)
	GetByID(id string) (models.Admin, error)
	Create(admin models.Admin) (models.Admin, error)
	Update(admin models.Admin) (models.Admin, error)
	Delete(admin models.Admin) error
}

type adminRepository struct{}

func NewAdminRepository() AdminRepository {
	return &adminRepository{}
}

func (r *adminRepository) GetAll() ([]models.Admin, error) {
	var admins []models.Admin
	err := db.DB.Find(&admins).Error
	return admins, err
}

func (r *adminRepository) GetByID(id string) (models.Admin, error) {
	var admin models.Admin
	err := db.DB.First(&admin, "id = ?", id).Error
	return admin, err
}

func (r *adminRepository) Create(admin models.Admin) (models.Admin, error) {
	err := db.DB.Create(&admin).Error
	return admin, err
}

func (r *adminRepository) Update(admin models.Admin) (models.Admin, error) {
	err := db.DB.Save(&admin).Error
	return admin, err
}

func (r *adminRepository) Delete(admin models.Admin) error {
	err := db.DB.Delete(&admin).Error
	return err
}
