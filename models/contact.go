package models

type Contact struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  FullName  string `json:"full_name"`
  Email string `json:"email"`
  PhoneNumbers []PhoneNumber `json:"phone_numbers" gorm:"foreignKey:contact_id"`
}