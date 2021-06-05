package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
  database, err := gorm.Open("sqlite3", "contacts.db")

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&Contact{})

  DB = database
}