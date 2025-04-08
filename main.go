// main.go
package main

import (
    "log"
    "time"
    "prueba/config"
    "prueba/models"
    "prueba/routes"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()

    if err := models.InitDB(); err != nil {
        log.Fatalf("Error al conectar con la base de datos: %v", err)
    }
    defer models.DB.Close()

    if err := models.CreateTables(); err != nil {
        log.Fatalf("Error al crear tablas: %v", err)
    }

    r := gin.Default()

    // Configurar CORS ANTES de definir rutas
    configureCORS(r)

    // Definir rutas
    routes.SetupRoutes(r)

    // Iniciar servidor
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}

func configureCORS(r *gin.Engine) {
    corsConfig := cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: false,
        MaxAge:           12 * time.Hour,
    }

    r.Use(cors.New(corsConfig))
}
