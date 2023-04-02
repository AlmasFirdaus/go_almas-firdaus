package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Id     int    `json:"id" form:"id"`
	IdUser User   `json:"iduser" form:"iduser" gorm:"foreignKey:Id;references:Id"`
	Judul  string `json:"judul" form:"judul"`
	Konten string `json:"konten" form:"konten"`
}
