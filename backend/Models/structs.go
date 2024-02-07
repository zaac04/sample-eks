package models

type Todo struct {
	Id               string `gorm:"primaryKey"`
	Todo         string 
}