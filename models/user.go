package models

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"type:varchar(100);unique" json:"username" binding:"required,min=3"`
	Password string `gorm:"type:varchar(100)" json:"password" binding:"required,min=7"`
	Name     string `gorm:"type:varchar(100)" json:"name" binding:"required,min=3"`
}
