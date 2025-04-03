package routes

import (
    "prueba/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    // Nota: El middleware CORS ya está configurado en main.go

    api := r.Group("/api")
    {
        api.POST("/persons", controllers.CreatePerson)
        api.GET("/persons", controllers.GetAllPersons)
        api.GET("/stats/genders", controllers.GetGenderCounts)
        api.GET("/stats/genders/short-polling", controllers.ShortPollingGenderCount) // Long Polling

    }

    return r
}