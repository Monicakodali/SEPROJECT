package models

import "time"

type Review struct {
	Email    string    `json:"Email"`
	Name     string    `json:"Name"`
	Est_id   string    `json:"Est_id"`
	Est_name string    `json:"Est_name"`
	Review   string    `json:"Review"`
	Rating   float32   `json:"Rating"`
	revTime  time.Time `json:"revTime"`
}
