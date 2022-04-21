package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Monicakodali/SEPROJECT/api/controller"
	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/Monicakodali/SEPROJECT/api/utils"
	"github.com/gin-gonic/gin"
)

func TestInsertReview(t *testing.T) {

	//check DB connection
	db, err := utils.GetDBInstance()
	if err != nil {
		t.Fatal("Couldn't connect to Database")
	}
	db.LogMode(true)
	db.Debug().AutoMigrate(&models.Review{})
	defer db.Close()

	router := gin.New()

	reviewController := controller.ReviewController{}
	reviewController.Init(db)

	router.POST("/api/establishments", reviewController.NewReview)

	var jsonStr = []byte(`{
		"Review_user": "monica",
    	"Review_est": 7,
    	"Review": "Good food",
    	"Rating":2
	  }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/establishments", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "applocation/json")
	router.ServeHTTP(w, req)

	var jsonStr2 = []byte(`{
		"Review_user": "cpagolu",
    	"Review_est": 2,
    	"Review": "Good Dining place!",
    	"Rating":4
	  }`)
	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/api/establishments", bytes.NewBuffer(jsonStr2))
	req1.Header.Set("Content-Type", "applocation/json")
	router.ServeHTTP(w1, req1)

}
func TestGetReview(t *testing.T) {

	//check DB connection
	db, err := utils.GetDBInstance()
	if err != nil {
		t.Fatal("Couldn't connect to Database")
	}
	db.LogMode(true)
	db.Debug().AutoMigrate(&models.Establishment{}, &models.Review{})
	defer db.Close()

	router := gin.New()

	reviewController := controller.ReviewController{}
	reviewController.Init(db)

	router.GET("/api/establishments", reviewController.GetReviewsForEst)

	var jsonStr = []byte(`{"est_id": 76}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/est/review", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("Establishment ID", string(jsonStr))
	router.ServeHTTP(w, req)

}

func TestRemoveReview(t *testing.T) {

	//check DB connection
	db, err := utils.GetDBInstance()
	if err != nil {
		t.Fatal("Couldn't connect to Database")
	}
	db.LogMode(true)
	db.Debug().AutoMigrate(&models.Establishment{}, &models.UFDining{})
	defer db.Close()

	router := gin.New()

	establishmentController := controller.EstController{}
	establishmentController.Init(db)

	router.POST("/api/establishments", establishmentController.GetOneEstHandler)
	router.GET("/api/establishments", establishmentController.GetOneEstHandler)
	router.DELETE("/api/establishments", establishmentController.GetOneEstHandler)

	var jsonStr = []byte(`{}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/establishments", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

}
