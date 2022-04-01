package models

//establishment structure
type Establishment struct {
	Est_Id       string  `gorm:"primaryKey;not null;unique" json:"est_id"`
	Type         string  `gorm:"not null" json:"Type"`
	Name         string  `json:"Name"`
	X_coordinate float64 `json:"x"`
	Y_coordinate float64 `json:"y"`
}
