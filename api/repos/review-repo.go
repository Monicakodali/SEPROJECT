package repos

import (
	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/jinzhu/gorm"
)

type RevRepo struct {
	db *gorm.DB
}

func (revRepo *RevRepo) Init(db *gorm.DB) {
	revRepo.db = db
}

func (revRepo *RevRepo) GetRevByUser(email string) ([]models.Review, error) {

	var reviews []models.Review

	query := revRepo.db.Where("Email = ?", email).Find(&reviews)
	if query.Error != nil {
		return reviews, query.Error
	}
	return reviews, nil
}

func (revRepo *RevRepo) GetRevByEst(ename string) ([]models.Review, error) {

	var reviews []models.Review

	query := revRepo.db.Where("Est_name = ?", ename).Find(&reviews)
	if query.Error != nil {
		return reviews, query.Error
	}
	return reviews, nil
}

func (revRepo *RevRepo) GetAllReviews() ([]models.Review, error) {

	var reviews []models.Review

	query := revRepo.db.Find(&reviews)
	if query.Error != nil {
		return nil, query.Error
	}
	return reviews, nil
}

func (revRepo *RevRepo) CreateRev(revab models.Review) error {

	query := revRepo.db.Create(&revab)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (revRepo *RevRepo) DeleteRev(email string, eid string) error {

	query := revRepo.db.Delete(&models.Establishment{}, email, eid)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
