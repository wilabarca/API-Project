package config

type Config struct {
    Database struct {
        Host     string
        Port     string
        User     string
        Password string
        Name     string
    }
}

var AppConfig Config

func LoadConfig() {
    AppConfig.Database.Host = "localhost"
    AppConfig.Database.Port = "3306"
    AppConfig.Database.User = "abarca"
    AppConfig.Database.Password = "1234abarca" // Cambia esto
    AppConfig.Database.Name = "persona"
}