// routes/routes.go
package routes

import (
    "prueba/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.POST("/persons", controllers.CreatePerson)
        api.GET("/persons", controllers.GetAllPersons)
        api.GET("/stats/genders", controllers.GetGenderCounts)
        api.GET("/stats/genders/short-polling", controllers.ShortPollingGenderCount)
    }
}
