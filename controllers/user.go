package controllers

import (
    "net/http"
    "gorest/models"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
    Email       string `json:"email" binding:"required"`
    Phone       string `json:"phone" binding:"required"`
    Password    string `json:"password" binding:"required"`
}

func GetUsers(c *gin.Context) {
    emailFilter := c.Query("email") // email url param

    query := models.DB

    if emailFilter != "" {
        query = query.Where("email = ?", emailFilter)
    }

    var users []models.User
    query.Find(&users)

    c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
    var input CreateUserInput 

    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
        return
    }

    user := models.User{
        Email: input.Email, 
        Phone: input.Phone,
        Password: string(hashedPassword),
    }

    models.DB.Create(&user)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
    var user models.User
    if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    models.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"data": "Succesfully deleted a user."})
}
