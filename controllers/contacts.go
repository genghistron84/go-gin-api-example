package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/genghistron84/go-gin-api-example/models"
)

// GET /contacts
func FindContacts(c *gin.Context) {
  var contacts []models.Contact
  models.DB.Find(&contacts)
  c.JSON(http.StatusOK, gin.H{"data": contacts})
}