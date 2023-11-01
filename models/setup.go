package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=fullstack-postgres user=admin password=admin dbname=gorest port=5432 sslmode=disable timezone=Europe/Amsterdam"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})  // change the database provider if necessary

    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&User{})  // register User model

    DB = database
}
