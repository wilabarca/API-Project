package routes

import (
    "prueba/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        api.POST("/persons", controllers.CreatePerson)
        api.GET("/persons", controllers.GetAllPersons)
        api.GET("/stats/genders", controllers.GetGenderCounts)
    }

    return r
}