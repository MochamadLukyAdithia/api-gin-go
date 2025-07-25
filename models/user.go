package models

type User struct {
    ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email" gorm:"unique"`
}
