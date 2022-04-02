package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Monicakodali/SEPROJECT/api/controller"
	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/Monicakodali/SEPROJECT/api/utils"
	"github.com/gin-gonic/gin"
)

func TestCreateReview(t *testing.T) {
	db, err := utils.GetDBInstance()
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	db.Debug().AutoMigrate(&models.Review{})
	defer db.Close()
	//fmt.Println(db)

	router := gin.New()

	reviewController := controller.ReviewController{}
	reviewController.Init(db)

	router.POST("/api/users", reviewController.NewReview)

	var jsonStr = []byte(`{
		"Review_id": 88,
    "Review_user": "cpagolu",
    "Review": "Great Service!",
    "Review_est": 82,
    "Rating": 4}`)

	w := httptest.NewRecorder()
	//b, _ := json.Marshal(jsonStr)
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var jsonStr2 = []byte(`{
		"Review_id": 82,
		"Review_user": "cpagolu",
		"Review": "Great Service!",
		"Review_est": 82,
		"Rating": 4}`)
	w1 := httptest.NewRecorder()
	b1, _ := json.Marshal(jsonStr2)
	req1, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(b1))
	req1.Header.Set("Content-Type", "application/json; charset=UTF-8")
	router.ServeHTTP(w1, req1)
}

func TestGetReview(t *testing.T) {

	//check DB connection
	db, err := utils.GetDBInstance()
	if err != nil {
		t.Fatal("Couldn't connect to Database")
	}
	db.LogMode(true)
	db.Debug().AutoMigrate(&models.User{})
	defer db.Close()

	router := gin.New()

	userController := controller.UserController{}
	userController.Init(db)

	router.GET("/api/users", userController.GetUser)

	var jsonStr = []byte(`{"Email": "mmisra@ufl.edu"}`)
	b, _ := json.Marshal(jsonStr)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Encode()
	fmt.Println(q)
	router.ServeHTTP(w, req)

}

func TestRemoveReview(t *testing.T) {

	//check DB connection
	db, err := utils.GetDBInstance()
	if err != nil {
		t.Fatal("Couldn't connect to Database")
	}
	db.LogMode(true)
	db.Debug().AutoMigrate(&models.User{})
	defer db.Close()

	router := gin.New()

	userController := controller.UserController{}
	userController.Init(db)

	router.POST("/api/establishments", userController.SignUp)
	router.GET("/api/establishments", userController.GetUser)
	router.DELETE("/api/establishments", userController.DeleteUser)

	var jsonStr = []byte(`{}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

}
