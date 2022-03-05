package controller

import (
	"net/http"

	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/Monicakodali/SEPROJECT/api/repos"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ReviewController struct {
	revRepo *repos.RevRepo
}

func (rInstance *ReviewController) Init(db *gorm.DB) {
	rInstance.revRepo = &repos.RevRepo{}
	rInstance.revRepo.Init(db)
}

func (rev *ReviewController) ListReviews(c *gin.Context) {

	res, err := rev.revRepo.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (rev *ReviewController) NewReview(ctx *gin.Context) {
	var rInstance models.Review

	if err := ctx.ShouldBindJSON(&rInstance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := rev.revRepo.AddReview(rInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, rInstance)
}
