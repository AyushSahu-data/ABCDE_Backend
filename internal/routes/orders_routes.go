package routes

import (
    "github.com/gin-gonic/gin"
    "Cart_Backend/internal/handlers"
)

func SetupOrderRoutes(router *gin.Engine) {
    router.GET("/orders", handlers.GetOrders)
}
