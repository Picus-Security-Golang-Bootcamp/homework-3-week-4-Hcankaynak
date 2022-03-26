package city

import "gorm.io/gorm"

type CityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{db: db}
}
