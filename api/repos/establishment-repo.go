package repos

import (
	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/jinzhu/gorm"
)

type EstRepo struct {
	db *gorm.DB
}

func (estRepo *EstRepo) Init(db *gorm.DB) {
	estRepo.db = db
}

func (estRepo *EstRepo) GetEstByID(eid string) (models.Establishment, error) {

	var establishment models.Establishment

	query := estRepo.db.Where("id = ?", eid).First(&establishment)
	if query.Error != nil {
		return establishment, query.Error
	}
	return establishment, nil
}

func (estRepo *EstRepo) GetAllEst() ([]models.Establishment, error) {

	var establishments []models.Establishment

	query := estRepo.db.Find(&establishments)
	if query.Error != nil {
		return nil, query.Error
	}
	return establishments, nil
}

func (estRepo *EstRepo) CreateEst(estab models.Establishment) error {

	query := estRepo.db.Create(&estab)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (estRepo *EstRepo) DeleteEst(eid string) error {

	query := estRepo.db.Delete(&models.Establishment{}, eid)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
