package repositories

import (
	"housybe/models"

	"gorm.io/gorm"
)

type HouseRepository interface {
	FindHouses() ([]models.House, error)
	GetHouse(ID int) (models.House, error)
	AddHouse(house models.House) (models.House, error)
	UpdateHouse(house models.House) (models.House, error)
	DeleteHouse(house models.House) (models.House, error)
}

func RepositoryHouse(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindHouses() ([]models.House, error) {
	var houses []models.House
	err := r.db.Find(&houses).Error

	return houses, err
}

func (r *repository) GetHouse(ID int) (models.House, error) {
	var house models.House
	err := r.db.First(&house, ID).Error

	return house, err
}

func (r *repository) AddHouse(house models.House) (models.House, error) {
	err := r.db.Create(&house).Error

	return house, err
}

func (r *repository) UpdateHouse(house models.House) (models.House, error) {
	err := r.db.Save(&house).Error

	return house, err
}

func (r *repository) DeleteHouse(house models.House) (models.House, error) {
	err := r.db.Delete(&house).Error

	return house, err
}
