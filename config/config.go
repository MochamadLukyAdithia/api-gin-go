package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var PORT string

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("❗️ .env tidak ditemukan, gunakan default env")

        PORT = "8080"
        return
    }

    PORT = os.Getenv("PORT")
    if PORT == "" {
        PORT = "8080"
    }
}
