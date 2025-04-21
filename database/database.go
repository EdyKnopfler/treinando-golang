package database

import (
	"embed"
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
