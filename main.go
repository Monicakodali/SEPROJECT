package main

import (
	"net/http"

	"github.com/Monicakodali/SEPROJECT/api/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db, err := gorm.Open("sqlite3", "seproj.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.Close()

	router := gin.New()

	establishmentController := controller.App{}
	establishmentController.Initialize("sqlite3", "seproj.db")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	api := router.Group("/api/establishments")
	{
		api.POST("/", establishmentController.createEstablishments)
	}

	router.Run()

}
