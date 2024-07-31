package models

type Bus struct {
	ID       uint        `json:"id" gorm:"primaryKey"`
	Number   uint        `json:"number"`
	Stations JSONIntList `json:"favorite_numbers" gorm:"type:json"`
}
