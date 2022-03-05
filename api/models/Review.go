package models

import (
	"time"
)

type Review struct {
	Email    string    `json:"Email"`
	Name     string    `json:"Name"`
	Est_id   string    `json:"Est_id"`
	Est_name string    `json:"Est_name"`
	Review   string    `json:"Review"`
	Rating   float64   `json:"Rating"`
	RevTime  time.Time `json:"revTime"`
}
