package main

import (
  "github.com/gin-gonic/gin"
  "github.com/genghistron84/go-gin-api-example/models"
  "github.com/genghistron84/go-gin-api-example/controllers"
)

func main() {
  r := gin.Default()
  models.ConnectDatabase()
  r.GET("/contacts", controllers.GetContacts)
  r.POST("/contact", controllers.CreateContact)
  r.GET("/contact/:id", controllers.GetContact)
  r.PATCH("/contact/:id", controllers.UpdateContact)
	r.DELETE("/contact/:id", controllers.DeleteContact)
  r.Run()
}