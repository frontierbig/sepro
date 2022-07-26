package main
 
import (
  "github.com/frontierbig/SEPRO/controller"
  "github.com/frontierbig/SEPRO/entity"
  "github.com/gin-gonic/gin"
)
 
func main() {
  entity.SetupDatabase()
 
  r := gin.Default()
 
  // User Routes
  r.GET("/users", controller.ListUserMember)
  r.GET("/user/:id", controller.GetUser)
  r.POST("/users", controller.CreateUser)
  r.PATCH("/users", controller.UpdateUser)
  r.DELETE("/users/:id", controller.DeleteUser)
  
  // Run the server
  r.Run()
}
