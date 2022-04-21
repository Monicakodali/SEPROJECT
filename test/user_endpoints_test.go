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

func TestSignup(t *testing.T) {

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

	router.POST("/api/establishments", establishmentController.CreateEstablishments)

	var jsonStr = []byte(`{
		"Est_Id": 7983,
		"Name": "P.O.D. Market",
		"X_coordinate": 29.64636443982178,
		"Y_coordinate": -82.34316087863382,
		"IsOpen": 1,
		"Building": "551",
    "Room": "173",
    "Url": "https://gatordining.campusdish.com/LocationsAndMenus/BeatyPODMarket"
	  }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/establishments", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "applocation/json")
	router.ServeHTTP(w, req)

	var jsonStr2 = []byte(`{
		"Est_Id": 76,
		"Name": "Boar's Head",
		"X_coordinate": 29.648720068154468,
		"Y_coordinate": -82.34128718387561,
		"IsOpen": 1,
		"Building": "556",
    	"Room": "104",
    	"Url": ""
	  }`)
	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/api/establishments", bytes.NewBuffer(jsonStr2))
	req1.Header.Set("Content-Type", "applocation/json")
	router.ServeHTTP(w1, req1)

}
func TestLogin(t *testing.T) {

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

	router.GET("/api/establishments", establishmentController.GetOneEstHandler)

	var jsonStr = []byte(`{"est_id": 76}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/establishments", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("Establishment ID", string(jsonStr))
	router.ServeHTTP(w, req)

}

func TestDeleteUser(t *testing.T) {

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
