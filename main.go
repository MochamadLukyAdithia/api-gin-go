package main

import (
    "api-gin/config"
    "api-gin/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    config.ConnectDatabase()
    r := gin.Default()
    routes.SetupRoutes(r)

    r.Run(":" + config.PORT)
}
