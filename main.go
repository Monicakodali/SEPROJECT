package main

import (
	"fmt"
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

	defer db.Close()

	if !db.HasTable(&models.Establishment{}) {
		fmt.Println("Table doesnot exist. Creating table establishment")
		db.Debug().AutoMigrate(&models.Establishment{})
	}

	if !db.HasTable(&models.UFDining{}) {
		fmt.Println("Table doesnot exist. Creating table UFDining")
		db.Debug().AutoMigrate(&models.UFDining{})
	}

	if !db.HasTable(&models.User{}) {
		fmt.Println("Table doesnot exist. Creating table User")
		db.Debug().AutoMigrate(&models.User{})
	}

	if !db.HasTable(&models.Review{}) {
		fmt.Println("Table doesnot exist. Creating table Review")
		db.Debug().AutoMigrate(&models.Review{})
	}

	//fmt.Println(db)

	router := gin.New()

	establishmentController := controller.EstController{}
	establishmentController.Init(db)
	userController := controller.UserController{}
	userController.Init(db)
	revController := controller.ReviewController{}
	revController.Init(db)

	router.Use(func(ctx *gin.Context) {

		if ctx.Request.Header["Content-Length"] != nil && ctx.Request.Header["Content-Length"][0] == "0" {
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
