package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/genghistron84/go-gin-api-example/models"
)

type CreateContactInput struct {
	FullName  string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// GET /contacts
func GetContacts(c *gin.Context) {
  var contacts []models.Contact
  models.DB.Find(&contacts)
  c.JSON(http.StatusOK, gin.H{"contacts": contacts})
}

// POST /contact
func CreateContact(c *gin.Context) {
  // Validate input
  var input CreateContactInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create contact
  contact := models.Contact{FullName: input.FullName, Email: input.Email}
  models.DB.Create(&contact)

  c.JSON(http.StatusOK, gin.H{"data": contact})
}

// GET /contact/:id
func GetContact(c *gin.Context) {
	// Get model if it exists
	var contact models.Contact
	if err := models.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"contact": contact})
}

// PATCH /contact/:id
func UpdateContact(c *gin.Context) {

}

// DELETE /contact/:id
func DeleteContact(c *gin.Context) {

}
