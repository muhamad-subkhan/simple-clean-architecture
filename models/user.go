package models

type User struct {
	Id   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}