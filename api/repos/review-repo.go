package repos

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/jinzhu/gorm"
)

type RevRepo struct {
	db *gorm.DB
}

func (revRepo *RevRepo) Init(db *gorm.DB) {
	revRepo.db = db
}

type MyTime time.Time

func (m *MyTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	return json.Unmarshal(data, (*time.Time)(m))
}

// Get all reviews
func (revRepo *RevRepo) GetAllReviews() ([]models.Review, error) {

	var reviewList []models.Review

	query := revRepo.db.Debug().Find(&reviewList)
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
	fmt.Println("ADDING.....")

	query := revRepo.db.Debug().Create(&newReview)
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
