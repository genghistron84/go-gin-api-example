package models

type Phone struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  ContactID  int `json:"contact_id"`
  Contact Contact
  Phone string `json:"phone_number"`
}