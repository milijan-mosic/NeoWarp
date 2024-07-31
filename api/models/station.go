package models

type JSONIntList []uint

type Station struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name"`
	Latitude  float32     `json:"latitude"`
	Longitude float32     `json:"longitude"`
	Buses     JSONIntList `json:"favorite_numbers" gorm:"type:json"`
}
