package main

import (
	"net/http"

	"github.com/Monicakodali/SEPROJECT/api/controller"
	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/Monicakodali/SEPROJECT/api/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db, err := utils.GetDBInstance()
	if err != nil {
		panic("failed to connect database")
	}

	db.Debug().AutoMigrate(&models.Establishment{})
	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Review{})
	defer db.Close()
	//fmt.Println(db)

	router := gin.New()

	establishmentController := controller.EstController{}
	establishmentController.Init(db)
	userController := controller.UserController{}
	userController.Init(db)
	revController := controller.ReviewController{}
	revController.Init(db)

	router.Use(func(ctx *gin.Context) {

		if ctx.Request.Header["Content-Length"][0] == "0" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Payload should not be empty"})
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
	})

	router.GET("/api/establishments", establishmentController.ListEstHandler)
	router.POST("/api/establishments", establishmentController.CreateEstablishments)
	router.DELETE("/api/establishments", establishmentController.DeleteEstablishment)
	router.GET("/api/users", userController.ListUsers)
	router.POST("/api/users", userController.SignUp)
	router.DELETE("/api/users", userController.DeleteUser)
	router.GET("/api/reviews", revController.ListReviews)
	router.POST("/api/reviews", revController.NewReview)
	router.Run()
	//running

}
