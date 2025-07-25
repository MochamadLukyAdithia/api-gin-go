package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var PORT string

func LoadEnv() {
    // Hanya load .env saat di local development
    _ = godotenv.Load() // abaikan error, karena di production biasanya tidak ada .env

    PORT = os.Getenv("PORT")
    if PORT == "" {
        log.Println("❗️ PORT tidak ditemukan di env, gunakan default :8080")
        PORT = "8080"
    } else {
        log.Println("✅ Menggunakan PORT dari environment:", PORT)
    }
}
