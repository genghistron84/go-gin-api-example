package models

import (
  "github.com/jinzhu/gorm"
)

type Contact struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  FullName  string `json:"full_name"`
  Email string `json:"email"`
}