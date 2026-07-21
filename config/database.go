package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *Config) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Gagal mendapatkan instance database: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)           
	sqlDB.SetMaxOpenConns(100)          
	sqlDB.SetConnMaxLifetime(time.Hour) 

	DB = db
	log.Println("Berhasil terhubung ke database PostgreSQL")

	RunMigrations()
}

func RunMigrations() {
	log.Println("AutoMigrate selesai dijalankan (saat ini belum ada model)")
}

func CloseDatabase() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err == nil {
			err = sqlDB.Close()
			if err != nil {
				log.Printf("Gagal menutup koneksi database: %v", err)
			} else {
				log.Println("Koneksi database berhasil ditutup")
			}
		}
	}
}
