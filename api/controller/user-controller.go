package controller

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"strconv"
	"strings"
	"time"

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

type Credentials struct {
	email        string
	date         time.Time
	verification string
	validity     time.Time
	user         utils.NewUser
}
type UserController struct {
	usrRepo *repos.UsrRepo
}

func (uInstance *UserController) Init(db *gorm.DB) {
	uInstance.usrRepo = &repos.UsrRepo{}
	uInstance.usrRepo.Init(db)
}

func (usr *UserController) Login(ctx *gin.Context) {
	var body models.User
	// if err := ctx.BindJSON(&body); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Invalid request body",
	// 	})
	// 	return
	// }
	res, err := usr.usrRepo.GetUser(body.Username, body.Password)
	// Validate form input
	if strings.Trim(body.Username, " ") == "" || strings.Trim(body.Password, " ") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}
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

// Send verification mail

func (usr *UserController) SignUp(ctx *gin.Context) {

	var uInstance models.User

	if err := ctx.ShouldBindJSON(&uInstance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	email := uInstance.Email
	// Perform a very basic email check. We'll send a validation email anyway.
	if !strings.Contains(email, "@") {
		ctx.JSON(http.StatusInternalServerError, "Email wrong")
		return
	}

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

	fmt.Printf("\nstatus for %v is %v", email, myJson.Status)
	ctx.JSON(http.StatusInternalServerError, "Email wrong")

	user_creds := models.User{
		Username:  uInstance.Username,
		Email:     uInstance.Email,
		Password:  utils.EncryptPassword(uInstance.Password),
		FirstName: uInstance.FirstName,
		LastName:  uInstance.LastName,
		Verified:  0,
	}

	err1 := usr.usrRepo.AddUser(user_creds)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err1.Error(),
		})
		return
	}
	//ctx.JSON(201, gin.H{"message": "New user account registered"})
	ctx.JSON(http.StatusCreated, user_creds)

	ml := utils.NewUser{}
	verificationID := rand.Intn(100000000)
	idCreated := time.Now()
	ml.SetVerificationID(strconv.Itoa(verificationID), idCreated)
	ml.SetState(utils.StateCreated)
	// Check what needs to be done now.
	//b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	from := "mishramanjari18@gmail.com"
	header := make(map[string]string)
	header["From"] = from
	header["To"] = ml.GetEmail()
	header["Subject"] = fmt.Sprintf("=? UTF-8? B?% s? =", b64.StdEncoding.EncodeToString([]byte("message header 2")))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"
	Body := "Please verify your email!"
	template := ""
	for k, v := range header {
		template += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	template += "\r\n" + b64.StdEncoding.EncodeToString([]byte(Body))
	host := "smtp.mail.com"
	password := "password"

	auth := smtp.PlainAuth("", email, password, host)
	fmt.Printf("\nSending verification email for new account")
	// Send notification email.
	data := Credentials{
		email:        user_creds.Email,
		date:         time.Now(),
		verification: strconv.Itoa(verificationID),
		validity:     idCreated.Add(3 * 24 * time.Hour),
		user:         ml,
	}
	err2 := smtp.SendMail(
		host+":25",
		auth,
		from,
		[]string{user_creds.Email},
		[]byte(template),
	)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"\nCould not send verification code!": err2.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, data.email)

	// Is the verification ID still valid?
	if data.validity.Before(time.Now()) {
		fmt.Printf("\nVerification ID for user %s (%s) expired: %s", ml.GetID(), verificationID)
		err2 := smtp.SendMail(
			host+":25",
			auth,
			email,
			[]string{user_creds.Email},
			[]byte(template),
		)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Could not send verification code!": err2.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, data.email)

		return
	}

	// User has been verified. Update status.
	ml.SetState(utils.StateVerified)
	ml.SetVerificationID("", time.Unix(0, 0)) // Invalidate verification ID.
	emailcheck := ctx.Param("Email")
	result, err := usr.usrRepo.GetUserEmail(emailcheck)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	fmt.Printf("User %s (%s) has been verified", ml.GetID(), ml.GetEmail())
	result.Verified = 1
	ctx.JSON(http.StatusOK, &result)

	// b64 := base64.NewEncoding('ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789')
	// host := “smtp.mail.com”
	// email := “mail1@mail.com”
	// password := “password”
	// toEmail := “mail2@mail.com”
	// from := mail.Address {sender, email}
	// to := mail.Address {receiver, to email}

	// auth := smtp.PlainAuth(
	//     “”,
	//     email,
	//     password,
	//     host,
	// )
	// err := smtp.SendMail(
	//     host+”:25″,
	//     auth,
	//     email,
	//     []string{to.Address},
	//     []byte(message),
	// )
	// if err != nil {
	//     panic(err)
	// }
	// fmt.Println("Email Sent Successfully!")
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
