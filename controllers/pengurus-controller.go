package controllers

import (
    "api-gin/config"
    "api-gin/models"
    "api-gin/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetPengurus(c *gin.Context) {
    var pengurus []models.Pengurus
    if err := config.DB.Find(&pengurus).Error; err != nil {
        utils.Error(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.Success(c, pengurus)
}

func CreatePengurus(c *gin.Context) {
    var pengurus models.Pengurus
    if err := c.ShouldBindJSON(&pengurus); err != nil {
        utils.Error(c, http.StatusBadRequest, err.Error())
        return
    }
    if err := config.DB.Create(&pengurus).Error; err != nil {
        utils.Error(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.Success(c, pengurus)
}
