package models

type User struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Verified string `json:"Verified"`
}
