// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"

// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// 	"golang.org/x/crypto/bcrypt"
// )

// type User struct {
// 	Username string `json:"Username"`
// 	Email    string `json:"Email"`
// 	Name     string `json:"Name"`
// 	Password string `json:"Password"`
// 	Verified string `json:"Verified"`
// }

// var db *gorm.DB

// type App struct {
// 	DB *gorm.DB
// }

// func (a *App) Initialize(dbDriver string, dbURI string) {
// 	db, err := gorm.Open(dbDriver, dbURI)
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	a.DB = db

// 	db.AutoMigrate(&User{})
// }

// func GeneratePassword(password string) string {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	hashedPassword := string(hash)

// 	return hashedPassword
// }

// // func CheckPassword(userPassword string, correctPassword []byte) bool {
// // 	if err := bcrypt.CompareHashAndPassword(correctPassword, []byte(userPassword)); err != nil {
// // 		return false
// // 	}

// // 	return true
// // }

// func (a *App) getUser(c *gin.Context) {

// 	var user User
// 	Username := c.Param("Username")
// 	Password := c.Param("Password")
// 	a.DB.Where("Username = ?", Username).First(&user)
// 	if Username != user.Username {
// 		return
// 	}
// 	if result := a.DB.Where("Password = ?", GeneratePassword(Password)).First(&user); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	//a.DB.Find(&restaurants)

// 	c.JSON(http.StatusOK, &user)
// }

// func (a *App) listUsers(c *gin.Context) {
// 	//db := c.MustGet("db").(*gorm.DB)
// 	var users []User

// 	if result := a.DB.Find(&users); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	a.DB.Find(&users)

// 	c.JSON(http.StatusOK, &users)
// }

// func (a *App) createUsers(c *gin.Context) {
// 	var res User
// 	res.Password = GeneratePassword(res.Password)
// 	// toEmail := res.Email
// 	// // sender data
// 	// from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
// 	// password := os.Getenv("SMTPpwd")   // ex: "ieiemcjdkejspqz"
// 	// // receiver address privided through toEmail argument
// 	// to := []string{toEmail}
// 	// // smtp - Simple Mail Transfer Protocol
// 	// host := "smtp.gmail.com"
// 	// port := "587"
// 	// address := host + ":" + port
// 	// // message
// 	// subject := "Subject: Email Verification\r\n\r\n"
// 	// message := []byte(subject)
// 	// // athentication data
// 	// // func PlainAuth(identity, username, password, host string) Auth
// 	// auth := smtp.PlainAuth("", from, password, host)
// 	// // send mail
// 	// // func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
// 	// fmt.Println("message:", string(message))
// 	// err := smtp.SendMail(address, auth, from, to, message)
// 	// //return err
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// ver := "True"
// 	// res.Verified = ver
// 	if err := c.ShouldBindJSON(&res); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	if result := a.DB.Create(&res); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, &res)
// }

// func (a *App) deleteUser(c *gin.Context) {
// 	Username := c.Param("Username")

// 	if result := db.Delete(&User{}, Username); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// }

// func main() {

// 	var err error
// 	db, err = gorm.Open("sqlite3", "seproj.db")

// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db.Close()
// 	a := &App{db}
// 	a.Initialize("sqlite3", "seproj.db")

// 	r := gin.New()
// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Hello World!",
// 		})
// 	})

// 	// Create API route group
// 	//r.POST("/users/", a.createUsers)
// 	api := r.Group("/api/users")
// 	{
// 		api.POST("/", a.createUsers)
// 		//api.GET("/", a.listUsers)
// 		//api.GET("/:id", a.getUser)
// 		//api.DELETE("/:id", a.deleteUser)
// 	}
// 	r.Run()
// }
