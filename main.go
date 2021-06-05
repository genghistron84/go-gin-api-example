package main

import (
  "github.com/gin-gonic/gin"
  "github.com/genghistron84/go-gin-api-example/models"
  "github.com/genghistron84/go-gin-api-example/controllers"
)

func main() {
  r := gin.Default()
  models.ConnectDatabase()
  r.GET("/contacts", controllers.FindContacts)
  r.POST("/contact", controllers.CreateContact)
  r.Run()
}