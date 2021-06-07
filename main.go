package main

import (
  "github.com/gin-gonic/gin"  
  "github.com/gin-gonic/gin/binding"
  "github.com/genghistron84/go-gin-api-example/models"
  "github.com/genghistron84/go-gin-api-example/controllers"
  "github.com/ttacon/libphonenumber"
  "github.com/go-playground/validator/v10"
)

var australianNumber validator.Func = func(fl validator.FieldLevel) bool {
  phone_number := fl.Field().String()
  num, err := libphonenumber.Parse(phone_number, "AU")
  if err != nil {        
    // parsing failed
		return false    
  }  
  
  if libphonenumber.IsValidNumberForRegion(num, "AU") {
    // valid number
    return true
  }
  // fallback
  return false
}

func main() {
  r := gin.Default()
  models.ConnectDatabase()
  
  // bind custom E164 Austrlian validation function
  if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("australianNumber", australianNumber)
	}
  
  // routes  
  r.GET("/contacts", controllers.GetContacts)
  r.POST("/contact", controllers.CreateContact)
  r.GET("/contact/:id", controllers.GetContact)
  r.GET("/contact/:id/numbers", controllers.GetContactNumbers)
  r.PATCH("/contact/:id", controllers.UpdateContact)
	r.DELETE("/contact/:id", controllers.DeleteContact)
  r.Run()
}

