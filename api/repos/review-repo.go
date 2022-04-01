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

func (revRepo *RevRepo) GetAllReviews() ([]models.Review, error) {

	var reviewList []models.Review

	query := revRepo.db.Find(&reviewList)
	if query.Error != nil {
		return nil, query.Error
	}
	return reviewList, nil
}

func (revRepo *RevRepo) GetReviewsForEst(eid string) ([]models.Review, error) {

	var reviewList []models.Review
	query := revRepo.db.Where("Est_id = ?", eid).Find(&reviewList)

	if query.Error != nil {
		return nil, query.Error
	}
	return reviewList, nil
}

func (revRepo *RevRepo) AddReview(newReview models.Review) error {

	query := revRepo.db.Create(&newReview)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
