package models

type User struct {
	Username  string `gorm:"primaryKey;not null;unique" json:"Username"`
	Email     string `gorm:"not null" json:"Email"`
	Password  string `gorm:"not null" json:"Password"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Address   string `json:"Address"`
	Verified  int    `gorm:"not null;default:0" json:"Verified"`
}
