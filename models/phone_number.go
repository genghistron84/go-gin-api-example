package models

type PhoneNumber struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  ContactID  int `json:"contact_id"`
  Phone string `json:"phone_number"  binding:"required,australianNumber"`
}