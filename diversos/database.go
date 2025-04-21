package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"com.derso/testify/database"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDatabase() (*gorm.DB, *sql.DB) {
	gdb, err := gorm.Open(sqlite.Open(database.DB_NAME), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := gdb.DB()

	if err != nil {
		log.Fatal(err)
	}

	return gdb, sqlDB
}

func runMigrations(sqlDB *sql.DB) {
	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatal(err)
	}

	goose.SetBaseFS(database.EmbedMigrations)
	path := "."

	if err := goose.Up(sqlDB, path); err != nil {
		log.Fatal(err)
	}
}

func main() {
	gdb, sqlDB := connectDatabase()
	defer sqlDB.Close()

	runMigrations(sqlDB)

	gdb = gdb.Debug()
	var zimTom database.Pet
	result := gdb.First(&zimTom, "nome = ?", "ZimTom")

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("1ª execução: registro não encontrado, foi inserido")
		gdb.Begin()
		gdb.Create(database.Pets)
		gdb.Commit()
	} else {
		fmt.Println(result)
		fmt.Println(zimTom.Nome, zimTom.Idade)
		gdb.Rollback()
	}

}
