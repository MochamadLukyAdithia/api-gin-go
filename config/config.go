package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

var PORT string

func LoadEnv() {
    _ = godotenv.Load() 
    PORT = os.Getenv("PORT")
    if PORT == "" {
        log.Println("❗️ PORT tidak ditemukan di env, gunakan default :8080")
        PORT = "8080"
    } else {
        log.Println("✅ Menggunakan PORT dari environment:", PORT)
    }
}
