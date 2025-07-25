package config

import (
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := os.Getenv("DATABASE_URL") // ambil dari .env

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Gagal konek ke PostgreSQL:", err)
    }

    DB = db
    log.Println("✅ Berhasil konek ke PostgreSQL")
}
