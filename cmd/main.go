package main

import (
    "github.com/gin-gonic/gin"
    "D:/Cart_Backend/internal/routes"
)

func main() {
    router := gin.Default()

    // Setup routes
    routes.SetupUserRoutes(router)
    routes.SetupItemRoutes(router)
    routes.SetupCartRoutes(router)
    routes.SetupOrderRoutes(router)

    // Run the server
    router.Run(":8080")
}
