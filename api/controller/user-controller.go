package controller

import (
	"net/http"

	"github.com/Monicakodali/SEPROJECT/api/models"
	"github.com/Monicakodali/SEPROJECT/api/repos"
	"github.com/Monicakodali/SEPROJECT/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UserController struct {
	usrRepo *repos.UsrRepo
}

func (uInstance *UserController) Init(db *gorm.DB) {
	uInstance.usrRepo = &repos.UsrRepo{}
	uInstance.usrRepo.Init(db)
}

func (usr *UserController) GetUser(ctx *gin.Context) {

	username := ctx.Param("Username")
	password := ctx.Param("Password")
	encrypted_password := utils.EncryptPassword(password)
	res, err := usr.usrRepo.GetUser(username, encrypted_password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid Credentials",
		})

		ctx.JSON(http.StatusOK, &res)
	}
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
	user_creds := models.User{
		Username: uInstance.Username,
		Email:    uInstance.Email,
		Name:     uInstance.Name,
		Password: utils.EncryptPassword(uInstance.Password),
		Verified: uInstance.Verified,
	}

	err := usr.usrRepo.AddUser(user_creds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

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
