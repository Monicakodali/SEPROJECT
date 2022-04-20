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

type Result struct {
	Est_Id       int
	Name         string
	X_coordinate float64
	Y_coordinate float64
	IsOpen       int
	Building     string
	Room         string
	Url          string
}

func (estRepo *EstRepo) GetEstByID(eid string) (Result, error) {
	var res Result
	//var establishments models.Establishment
	est_id, err := strconv.Atoi(eid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(est_id)
	query := estRepo.db.Debug().Raw("SELECT e.est_id, e.name, e.x_coordinate, e.y_coordinate, d.building, d.room, d.url FROM establishments e INNER JOIN uf_dinings d ON  e.est_id = d.diner_id WHERE e.est_id = ?", est_id).Scan(&res)
	if query.Error != nil {
		return res, query.Error
	}
	return res, nil
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

	est_id, err := strconv.Atoi(eid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(est_id)
	query := estRepo.db.Where("est_id = ?", est_id).Delete(&models.Establishment{})
	if query.Error != nil {
		return query.Error
	}
	return nil
}
