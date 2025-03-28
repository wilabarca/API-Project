package models

import (
    "database/sql"
    "fmt"
    "prueba/config"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
    cfg := config.AppConfig.Database
    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

    var err error
    DB, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        return err
    }

    return DB.Ping()
}

func CreateTables() error {
    query := `
    CREATE TABLE IF NOT EXISTS persons (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        edad INT NOT NULL,
        genero VARCHAR(50) NOT NULL,
        sexo VARCHAR(50) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    _, err := DB.Exec(query)
    return err
}