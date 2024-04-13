package routes

import (
    "github.com/gin-gonic/gin"
    "Cart_Backend/internal/handlers"
)

func SetupItemRoutes(router *gin.Engine) {
    router.POST("/items", handlers.CreateItem)
}
