package routes

import (
    "github.com/gin-gonic/gin"
    "Cart_Backend/internal/handlers"
)

func SetupCartRoutes(router *gin.Engine) {
    router.POST("/carts", handlers.CreateCart)
    router.POST("/carts/orders", handlers.ConvertToOrder)
}
