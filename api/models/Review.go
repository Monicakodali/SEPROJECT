package models

import "time"

type Review struct {
	Review_Id   int       `gorm:"primaryKey;autoIncrement:true;" json:"Review_id"`
	Review_user string    `json:"Review_user"`
	Review_est  int       `json:"Review_est"`
	Review      string    `json:"Review"`
	Rating      float64   `gorm:"not null" json:"Rating"`
	RevTime     time.Time `gorm:"type:datetime;autoCreateTime:milli" json:"revTime"`
}
