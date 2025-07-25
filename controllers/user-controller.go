package controllers

import (
    "api-gin/config"
    "api-gin/models"
    "api-gin/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetUsers(c *gin.Context) {
    var users []models.User
    config.DB.Find(&users)
    utils.Success(c, users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.Error(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := config.DB.Create(&user).Error; err != nil {
        utils.Error(c, http.StatusInternalServerError, err.Error())
        return
    }

    utils.Success(c, user)
}
