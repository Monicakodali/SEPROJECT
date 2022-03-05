package controller

import (
	"net/http"

	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/Monicakodali/SEPROJECT/api/repos"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type RevController struct {
	revRepo *repos.RevRepo
}

func (rRepo *RevController) Init(db *gorm.DB) {
	rRepo.revRepo = &repos.RevRepo{}
	rRepo.revRepo.Init(db)
}

func (rev *RevController) GetRatingHandler(ctx *gin.Context) {

	email := ctx.Param("Email")
	res, err := rev.revRepo.GetRevByUser(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		ctx.JSON(http.StatusOK, &res)
	}
}

func (rev *RevController) GetRaterHandler(ctx *gin.Context) {

	ename := ctx.Param("Est_name")
	res, err := rev.revRepo.GetRevByEst(ename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		ctx.JSON(http.StatusOK, &res)
	}
}

func (rev *RevController) ListRevHandler(c *gin.Context) {

	res, err := rev.revRepo.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (rev *RevController) CreateReviews(ctx *gin.Context) {

	var rInstance models.Review

	if err := ctx.ShouldBindJSON(&rInstance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := rev.revRepo.CreateRev(rInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, rInstance)
}

func (rev *RevController) DeleteReview(ctx *gin.Context) {
	eid := ctx.Param("Est_id")
	email := ctx.Param("Email")
	err := rev.revRepo.DeleteRev(email, eid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
