package main

import (
    "log"
    "prueba/config"
    "prueba/models"
    "prueba/routes"
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

    // Iniciar servidor
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}