package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"os"

	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/Monicakodali/SEPROJECT/api/repos"
	"github.com/Monicakodali/SEPROJECT/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type RealEmailResponse struct {
	Status string `json:"status"`
}

type UserController struct {
	usrRepo *repos.UsrRepo
}

var htmlBody = `
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<title>Hello!</title>
</head>
<body>
	<p> Please verify your email </p>
</body>
`

func (uInstance *UserController) Init(db *gorm.DB) {
	uInstance.usrRepo = &repos.UsrRepo{}
	uInstance.usrRepo.Init(db)
}

func (usr *UserController) GetUser(ctx *gin.Context) {
	var body models.User
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	res, err := usr.usrRepo.GetUser(body.Username, body.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid Credentials",
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (usr *UserController) GetUserByEmail(ctx *gin.Context) {

	email := ctx.Param("Email")
	res, err := usr.usrRepo.GetUserEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, &res)
}

func (usr *UserController) ListUsers(c *gin.Context) {

	res, err := usr.usrRepo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (usr *UserController) SignUp(ctx *gin.Context) {

	var uInstance models.User

	if err := ctx.ShouldBindJSON(&uInstance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	email := uInstance.Email
	url := "https://isitarealemail.com/api/email/validate?email=" + url.QueryEscape(email)
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	var myJson RealEmailResponse
	json.Unmarshal(body, &myJson)

	fmt.Printf("status for %v is %v", email, myJson.Status)

	//email := uInstance.Email
	result, _ := usr.usrRepo.GetUserEmail(email)
	if result.Email != "" {
		ctx.JSON(403, gin.H{"message": "Email is already in use"})
		ctx.Abort()
		return
	}

	user_creds := models.User{
		Username: uInstance.Username,
		Email:    uInstance.Email,
		Name:     uInstance.Name,
		Password: utils.EncryptPassword(uInstance.Password),
		Verified: uInstance.Verified,
	}

	from := "mishramanjari18@gmail.com"
	to := []string{
		user_creds.Email,
	}
	//user := "9c1d45eaf7af5b"
	password := os.Getenv("PASSWD")
	host := "smtp.gmail.com"
	port := "587"
	msg := []byte("From: mishramanjari18@gmail.com\r\n" +
		"To: mmisra@ufl.edu\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")
	auth := smtp.PlainAuth("", from, password, host)

	err1 := smtp.SendMail(host+":"+port, auth, from, to, msg)

	if err1 != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
	user_creds.Verified = "1"
	error := usr.usrRepo.AddUser(user_creds)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": error.Error(),
		})
		return
	}
	ctx.JSON(201, gin.H{"message": "New user account registered"})
	ctx.JSON(http.StatusCreated, user_creds)
}

func (usr *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("Username")
	err := usr.usrRepo.RemoveUser(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
