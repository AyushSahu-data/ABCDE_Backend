package routes

import (
    "github.com/gin-gonic/gin"
    "Cart_Backend/internal/handlers"
)

func SetupUserRoutes(router *gin.Engine) {
    router.POST("/users", handlers.CreateUser)
    router.POST("/users/login", handlers.Login)
}
