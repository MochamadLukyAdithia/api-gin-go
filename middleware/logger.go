package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Printf("📥 %s %s\n", c.Request.Method, c.Request.URL.Path)
        c.Next()
        fmt.Printf("📤 Status: %d\n", c.Writer.Status())
    }
}
