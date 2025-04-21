package database

import (
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Pet struct {
	ID    uint   `gorm:"primaryKey"`
	Nome  string `gorm:"size:20;not null"`
	Idade uint   `gorm:"not null"`
}

const DB_NAME = "pets.db"

var Pets = []Pet{
	{
		ID:    1,
		Nome:  "ZimTom",
		Idade: 9,
	},
	{
		ID:    2,
		Nome:  "Dóia",
		Idade: 3,
	},
}

//go:embed *.sql
var EmbedMigrations embed.FS

func ConnectDatabase() (*gorm.DB, *sql.DB) {
	gdb, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := gdb.DB()

	if err != nil {
		log.Fatal(err)
	}

	return gdb, sqlDB
}

func RunMigrations(sqlDB *sql.DB) {
	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatal(err)
	}

	goose.SetBaseFS(EmbedMigrations)
	path := "."

	if err := goose.Up(sqlDB, path); err != nil {
		log.Fatal(err)
	}
}
