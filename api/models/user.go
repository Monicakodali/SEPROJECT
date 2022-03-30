package models

type User struct {
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Address   string `json:"Address"`
	Verified  string `json:"Verified"`
}
