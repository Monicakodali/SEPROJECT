package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Review struct {
	Email    string    `json:"Email"`
	Name     string    `json:"Name"`
	Est_id   string    `json:"Est_id"`
	Est_name string    `json:"Est_name"`
	Review   string    `json:"Review"`
	Rating   string    `json:"Rating"`
	revTime  time.Time `json:"revTime"`
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

	db.AutoMigrate(&Review{})
}

func (a *App) ReviewsByEID(c *gin.Context) {

	var rev_place []Review
	Place := c.Param("Est_name")
	a.DB.Where("Est_name = ?", Place).First(&rev_place)
	// if Place != rev_place.Review {
	// 	return
	// }
	if result := a.DB.Where("Est_name = ?").First(&rev_place); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	//a.DB.Find(&restaurants)

	c.JSON(http.StatusOK, &rev_place)
}

func (a *App) ReviewsbyID(c *gin.Context) {

	var rev_place []Review
	Name := c.Param("Name")
	a.DB.Where("Est_name = ?", Name).First(&rev_place)
	// if Place != rev_place.Review {
	// 	return
	// }
	if result := a.DB.Where("Est_name = ?").First(&rev_place); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	//a.DB.Find(&restaurants)

	c.JSON(http.StatusOK, &rev_place)
}

func (a *App) listReviews(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)
	var reviews []Review

	if result := a.DB.Find(&reviews); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	a.DB.Find(&reviews)

	c.JSON(http.StatusOK, &reviews)
}

func (a *App) createReviews(c *gin.Context) {
	var res Review

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

func (a *App) deleteReview(c *gin.Context) {
	Name := c.Param("Name")
	Place := c.Param("Est_name")
	if result := db.Delete(&Review{}, Name, Place); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func main() {

	var err error
	db, err = gorm.Open("sqlite3", "./seproj.db")

	if err != nil {
		panic("failed to connect database")
	}

	a := &App{db}
	a.Initialize("sqlite3", "./seproj.db")

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Create API route group
	//r.POST("/users/", a.createUsers)
	api := r.Group("/api/users")
	{
		api.POST("/", a.createReviews)
		//api.GET("/", a.listUsers)
		//api.GET("/:id", a.getUser)
		//api.DELETE("/:id", a.deleteUser)
	}
	r.Run()
}
