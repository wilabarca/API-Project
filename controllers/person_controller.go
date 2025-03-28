package controllers

import (
    "net/http"
    "prueba/models"

    "github.com/gin-gonic/gin"
)

type PersonInput struct {
    Nombre string `json:"nombre" binding:"required"`
    Edad   int    `json:"edad" binding:"required"`
    Genero string `json:"genero" binding:"required"`
    Sexo   string `json:"sexo" binding:"required"`
}

func CreatePerson(c *gin.Context) {
    var input PersonInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    person := models.Person{
        Nombre: input.Nombre,
        Edad:   input.Edad,
        Genero: input.Genero,
        Sexo:   input.Sexo,
    }

    if err := person.Create(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, person)
}

func GetAllPersons(c *gin.Context) {
    persons, err := models.GetAllPersons()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, persons)
}