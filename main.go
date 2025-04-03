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
    // Cargar configuración
    config.LoadConfig()

    // Inicializar base de datos
    if err := models.InitDB(); err != nil {
        log.Fatalf("Error al conectar con la base de datos: %v", err)
    }
    defer models.DB.Close()

    // Crear tablas si no existen
    if err := models.CreateTables(); err != nil {
        log.Fatalf("Error al crear tablas: %v", err)
    }

    // Configurar rutas
    r := routes.SetupRoutes()

    // Configurar CORS
    configureCORS(r)

    // Iniciar servidor
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}

func configureCORS(r *gin.Engine) {
    config := cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // Permitir solo el frontend en esta URL
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}, // Métodos permitidos
        AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"}, // Cabeceras permitidas
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true, // Permitir credenciales
        MaxAge:           12 * time.Hour, // Tiempo máximo para que las pre-solicitudes CORS sean almacenadas en caché
    }

    r.Use(cors.New(config))

    // Opcional: Manejo explícito de OPTIONS para las rutas
    r.OPTIONS("/api/persons", func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Status(204) // No Content
    })
    r.OPTIONS("/api/persons/", func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Status(204) // No Content
    })
}
