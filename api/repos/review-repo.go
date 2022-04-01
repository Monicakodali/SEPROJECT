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

// Get all reviews
func (revRepo *RevRepo) GetAllReviews() ([]models.Review, error) {

	var reviewList []models.Review

	query := revRepo.db.Find(&reviewList)
	if query.Error != nil {
		return nil, query.Error
	}
	return reviewList, nil
}

// Get all the reviews for an establishment
func (revRepo *RevRepo) GetReviewsForEst(est_id string) ([]models.Review, error) {

	var reviewList []models.Review
	query := revRepo.db.Where("Review_est = ?", est_id).Find(&reviewList)

	if query.Error != nil {
		return nil, query.Error
	}
	return reviewList, nil
}

// Get all the reviews given by a user
func (revRepo *RevRepo) GetReviewsForUser(username string) ([]models.Review, error) {

	var reviewList []models.Review
	query := revRepo.db.Where("Review_user = ?", username).Find(&reviewList)

	if query.Error != nil {
		return nil, query.Error
	}
	return reviewList, nil
}

// Add reviews into the table
func (revRepo *RevRepo) AddReview(newReview models.Review) error {

	query := revRepo.db.Create(&newReview)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// Delete a review
func (revRepo *RevRepo) DeleteReview(rev_id string) error {

	query := revRepo.db.Where("Review_id = ?", rev_id).Delete(&models.Review{})
	if query.Error != nil {
		return query.Error
	}
	return nil
}
