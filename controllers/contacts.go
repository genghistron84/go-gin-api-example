package controllers

import (  
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/ttacon/libphonenumber"
  "github.com/genghistron84/go-gin-api-example/models"
)

// validation struct
type CreateContactInput struct {
  FullName  string `json:"full_name" binding:"required"`
  Email string `json:"email"`
  PhoneNumbers []models.PhoneNumber `json:"phone_numbers" binding:"required,dive"`
}

// GET /contacts
func GetContacts(c *gin.Context) {
  // initialize array of contact objects
  var contacts []models.Contact
  // use gorm to return a list of all contacts
  models.DB.Find(&contacts)
  // convert to JSON and return
  c.JSON(http.StatusOK, gin.H{"contacts": contacts})
}

// POST /contact
func CreateContact(c *gin.Context) {
  // bind posted values and run validations  
  var input CreateContactInput
  if err := c.ShouldBindJSON(&input); err != nil {
    // return error message if validations fail
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  
  // iterate through phone numbers and convert to E164  
  for i := range input.PhoneNumbers {
    num, err := libphonenumber.Parse(input.PhoneNumbers[i].Phone, "AU")
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "invalid phone number"})
      return
    }
    formattedNum := libphonenumber.Format(num, libphonenumber.E164)
    input.PhoneNumbers[i].Phone = formattedNum
  }  

  // create contact object if validations pass
  contact := models.Contact{FullName: input.FullName, Email: input.Email, PhoneNumbers: input.PhoneNumbers}
  // persist to database
  models.DB.Create(&contact)
  // return JSON response
  c.JSON(http.StatusOK, gin.H{"data": contact})
}

// GET /contact/:id
func GetContact(c *gin.Context) {
	// initialize contact object
  var contact models.Contact
  // use eager loading to fetch contact record and associated phone number records
	if err := models.DB.Preload("PhoneNumbers").Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
    // return error message if contact not found
		c.JSON(http.StatusBadRequest, gin.H{"error": "contact not found!"})
		return
	}
  // return contact if found
	c.JSON(http.StatusOK, gin.H{"contact": contact})
}

// PATCH /contact/:id
func UpdateContact(c *gin.Context) {

}

// DELETE /contact/:id
func DeleteContact(c *gin.Context) {
  var contact models.Contact
  // check records exists before attempting to delete
  if err := models.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
  models.DB.Delete(&contact)
  c.JSON(http.StatusOK, gin.H{"data": "contact deleted"})
}

// GET /contacts/:id/numbers
func GetContactNumbers(c *gin.Context) {
  var phone_numbers []models.PhoneNumber
  if err := models.DB.Find(&phone_numbers).Where("contact_id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "contact not found!"})
		return
	}
  c.JSON(http.StatusOK, gin.H{"phone_numbers": phone_numbers})
}