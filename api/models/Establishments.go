package models

//establishment structure
type Establishment struct {
	Est_Id       int     `gorm:"primary_key;not null;unique" json:"est_id"`
	Type         string  `gorm:"not null" json:"Type"`
	Name         string  `json:"Name"`
	X_coordinate float64 `json:"x"`
	Y_coordinate float64 `json:"y"`
	IsOpen       int     `gorm:"default:0" json:"isOpen"`
}
