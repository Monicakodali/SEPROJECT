package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Establishment struct {
	ID           string  `json:"id"`
	TYPE         string  `json:"type"`
	NAME         string  `json:"name"`
	BUILDING     string  `json:"building"`
	ROOM         string  `json:"room"`
	URL          string  `json:"url"`
	X_COORDINATE float64 `json:"x"`
	Y_COORDINATE float64 `json:"y"`
}

var db *gorm.DB

type App struct {
	DB *gorm.DB
}

func (a *App) Initialize(dbDriver string, dbURI string) {
	db, err := gorm.Open(dbDriver, dbURI)
	if err != nil {
		panic("failed to connect database")
	}
	a.DB = db

	db.AutoMigrate(&Establishment{})
}

func (a *App) listEstHandler(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)
	var establishments []Establishment

	if result := a.DB.Find(&establishments); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	//a.DB.Find(&restaurants)

	c.JSON(http.StatusOK, &establishments)
}

func (a *App) createEstablishments(c *gin.Context) {
	var res []Establishment

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

func main() {

	var err error
	db, err = gorm.Open("sqlite3", "seproj.db")

	if err != nil {
		panic("failed to connect database")
	}

	a := &App{db}
	a.Initialize("sqlite3", "seproj.db")

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Create API route group
	api := r.Group("/establishments")
	{
		api.POST("/", a.createEstablishments)
		api.GET("/", a.listEstHandler)
		//api.DELETE("/:id", a.deleteEstablishment)
	}
	//r.POST("/establishments", a.createEstablishments)
	//r.GET("/establishments", a.listEstHandler)

	r.Run()
}
