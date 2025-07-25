package routes

import (
    "api-gin/controllers"
    "api-gin/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.Use(middleware.CustomLogger())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    user := r.Group("/users")
    {
        user.GET("/", controllers.GetUsers)
        user.POST("/", controllers.CreateUser)
    }

    pengurus := r.Group("/pengurus")
    {
        pengurus.GET("/", controllers.GetPengurus)
        pengurus.POST("/", controllers.CreatePengurus)
    }
}
