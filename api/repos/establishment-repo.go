package repos

import (
	"fmt"
	"strconv"

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
	est_id, _ := strconv.Atoi(eid)
	fmt.Println(est_id)
	estRepo.db.Preloads("uf_dinings")
	//query := estRepo.db.Where("est_id = ?", eid).Find(&establishment)
	query := estRepo.db.Table("establishments").Select("establishments.est_id,establishments.type, establishments.name, uf_dinings.building, uf_dinings.room, establishments.x_coordinate, establishments.y_coordinate").Joins("JOIN uf_dinings on establishments.est_id = uf_dinings.diner_id").Where("establishments.est_id = ?", eid).Find(&establishment)
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
