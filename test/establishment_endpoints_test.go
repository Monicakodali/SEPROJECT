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

func TestInsertEstablishment(t *testing.T) {

	//check DB connection
	db, err := utils.GetDBInstance()
	if err != nil {
		panic("failed to connect database")
	}
	db.Debug().Exec("PRAGMA foreign_keys = ON")

	db.Debug().AutoMigrate(&models.Establishment{}, &models.UFDining{})
	db.Debug().Model(&models.Establishment{}).Related(&models.UFDining{}, "Diner_Id")
	db.LogMode(true)
	defer db.Close()

	router := gin.New()

	establishmentController := controller.EstController{}
	establishmentController.Init(db)

	router.POST("/api/establishments", establishmentController.CreateEstablishments)
	var jsonStr = []byte(`{
		"est_id": 82,
    	"type": "DINING",
    	"name": "P.O.D. Market",
    	"x": 29.64636443982178,
    	"y": -82.34316087863382,
		"is_open": 1,
		"image_id": ''
		}
	  }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/establishments", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "applocation/json")
	router.ServeHTTP(w, req)
	var jsonStr2 = []byte(`{
		"est_id": 76,
    	"type": "DINING",
    	"name": "Boar's Head",
    	"x": 29.648720068154468,
    	"y": -82.34128718387561,
		"is_open": 1,
		"image_id": ''
		}
	  }`)
	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/api/establishments", bytes.NewBuffer(jsonStr2))
	req1.Header.Set("Content-Type", "applocation/json")
	router.ServeHTTP(w1, req1)

}
func TestGetEstablishment(t *testing.T) {

	db, err := utils.GetDBInstance()
	if err != nil {
		panic("failed to connect database")
	}
	db.Debug().Exec("PRAGMA foreign_keys = ON")

	db.Debug().AutoMigrate(&models.Establishment{}, &models.UFDining{})
	db.Debug().Model(&models.Establishment{}).Related(&models.UFDining{}, "Diner_Id")
	db.LogMode(true)
	defer db.Close()

	router := gin.New()

	establishmentController := controller.EstController{}
	establishmentController.Init(db)

	router.GET("/api/establishments", establishmentController.ListEstHandler)

	//var jsonStr = []byte(`{"est_id": 76}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/establishments", nil)
	req.Header.Set("Content-Type", "application/json")
	//q := req.URL.Query()
	//q.Add("Establishment ID", string(jsonStr))
	router.ServeHTTP(w, req)

}

func TestRemoveEstablishment(t *testing.T) {

	//check DB connection
	db, err := utils.GetDBInstance()
	if err != nil {
		panic("failed to connect database")
	}
	db.Debug().Exec("PRAGMA foreign_keys = ON")

	db.Debug().AutoMigrate(&models.Establishment{}, &models.UFDining{})
	db.Debug().Model(&models.Establishment{}).Related(&models.UFDining{}, "Diner_Id")
	db.LogMode(true)
	defer db.Close()

	router := gin.New()

	establishmentController := controller.EstController{}
	establishmentController.Init(db)

	router.DELETE("/api/establishments", establishmentController.GetOneEstHandler)

	var jsonStr = []byte(`{"est_id": 76}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/establishments", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

}
