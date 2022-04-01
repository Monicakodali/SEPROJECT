package models

//UFDining structure belongs to Establishment structure
type UFDining struct {
	Diner_Id string        `gorm:"primaryKey;not null" json:"Diner_id"`
	Building string        `json:"Building"`
	Room     string        `json:"Room"`
	Url      string        `json:"Url"`
	Est_Id   Establishment `gorm:"foreignKey:Diner_Id;references:Est_Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
