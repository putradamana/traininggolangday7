package models

import(
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"Email" json="email" form:"Email"`
	Name     string `json:"Name" json="email" form:"Name"`
	Password string `json:"Password" json="email" form:"Password"`
}