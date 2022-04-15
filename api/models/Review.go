package models

import (
	"time"
)

type Review struct {
	Review_Id   int           `gorm:"primaryKey; not null; unique; autoIncrement;" json:"Review_id"`
	Review_user string        `json:"Review_user"`
	Review_est  int           `json:"Review_est"`
	Review      string        `json:"Review"`
	Rating      float64       `gorm:"not null" json:"Rating"`
	RevTime     time.Time     `json:"revTime"`
	Username    User          `gorm:"foreignKey:Review_user;references:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Username"`
	Est_id      Establishment `gorm:"foreignKey:Review_est;references:Est_Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"est_id"`
}
