package controllers

import (
    "net/http"
    "prueba/models"
    "time"

    "github.com/gin-gonic/gin"
)
// Short Polling: Se consulta cada 5 segundos
func ShortPollingGenderCount(c *gin.Context) {
    counts, err := models.GetGenderCounts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"gender_counts": counts})
}
func GetGenderCounts(c *gin.Context) {
    // Configuración de long polling (30 segundos máximo)
    timeout := time.After(30 * time.Second)
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    initialCounts, err := models.GetGenderCounts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    for {
        select {
        case <-timeout:
            c.JSON(http.StatusOK, initialCounts)
            return
        case <-ticker.C:
            currentCounts, err := models.GetGenderCounts()
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }

            // Comparar si hay cambios
            changed := false
            if len(currentCounts) != len(initialCounts) {
                changed = true
            } else {
                for k, v := range currentCounts {
                    if initialCounts[k] != v {
                        changed = true
                        break
                    }
                }
            }

            if changed {
                c.JSON(http.StatusOK, currentCounts)
                return
            }
        }
    }
}

// Función auxiliar para comparar los conteos de género
func compareCounts(initial, current map[string]int) bool {
    if len(initial) != len(current) {
        return true
    }

    for key, value := range initial {
        if current[key] != value {
            return true
        }
    }
    return false
}