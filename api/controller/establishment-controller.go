package controller

import (
	"net/http"

	"github.com/MonicaKodali/SEPROJECT/api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type App struct {
	DB *gorm.DB
}

func (r *App) Init(DB *gorm.DB) {
	r.DB = DB
}

func (a *App) Initialize(dbDriver string, dbURI string) {
	db, err := gorm.Open(dbDriver, dbURI)
	if err != nil {
		panic("failed to connect database")
	}
	a.DB = db

	db.AutoMigrate(&models.Establishment{})
}

func (a *App) GetOneEstHandler(c *gin.Context) {

	var establishment models.Establishment
	id := c.Param("id")

	if result := a.DB.Where("id = ?", id).First(&establishment); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	//a.DB.Find(&restaurants)

	c.JSON(http.StatusOK, &establishment)
}

func (a *App) ListEstHandler(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)
	var establishments []models.Establishment

	if result := a.DB.Find(&establishments); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	//a.DB.Find(&restaurants)

	c.JSON(http.StatusOK, &establishments)
}

func (a *App) CreateEstablishments(c *gin.Context) {
	var res models.Establishment

	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := a.DB.Create(&res); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &res)
}

func (a *App) DeleteEstablishment(c *gin.Context) {
	id := c.Param("id")

	if result := db.Delete(&models.Establishment{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
