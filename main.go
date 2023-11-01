package main

import (
    "gorest/models"
    "gorest/controllers"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

     //connect to DB
    models.ConnectDatabase()

    //routes
    router.GET("/users", controllers.GetUsers)
    router.POST("/users", controllers.CreateUser)
    router.DELETE("/users/:id", controllers.DeleteUser)

    router.Run(":8080")
}
