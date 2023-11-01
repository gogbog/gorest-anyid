package models

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Email     string    `json:"email"`
    Phone   string    `json:"phone"`
    Password   string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
}
