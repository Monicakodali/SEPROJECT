package controller

import (
	"fmt"
	"net/http"

	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/Monicakodali/SEPROJECT/api/repos"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type EstController struct {
	estRepo *repos.EstRepo
}

func (eRepo *EstController) Init(db *gorm.DB) {
	eRepo.estRepo = &repos.EstRepo{}
	eRepo.estRepo.Init(db)
}

func (est *EstController) GetOneEstHandler(ctx *gin.Context) {

	eid := ctx.Param("id")
	res, err := est.estRepo.GetEstByID(eid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		ctx.JSON(http.StatusOK, &res)
	}
}

func (est *EstController) ListEstHandler(c *gin.Context) {

	res, err := est.estRepo.GetAllEst()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (est *EstController) CreateEstablishments(ctx *gin.Context) {

	var eInstance models.Establishment

	if err := ctx.ShouldBindJSON(&eInstance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(eInstance.ID)

	/*new_est := models.Establishment{
		ID:           eInstance.ID,
		TYPE:         eInstance.TYPE,
		NAME:         eInstance.NAME,
		BUILDING:     eInstance.BUILDING,
		URL:          eInstance.URL,
		X_COORDINATE: eInstance.X_COORDINATE,
		Y_COORDINATE: eInstance.Y_COORDINATE,
	}*/

	err := est.estRepo.CreateEst(&eInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, eInstance)
}

func (est *EstController) DeleteEstablishment(ctx *gin.Context) {
	eid := ctx.Param("id")
	err := est.estRepo.DeleteEst(eid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
