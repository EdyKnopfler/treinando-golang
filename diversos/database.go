package main

import (
	"errors"
	"fmt"

	"com.derso/testify/database"
	"gorm.io/gorm"
)

func main() {
	gdb, sqlDB := database.ConnectDatabase()
	defer sqlDB.Close()

	database.RunMigrations(sqlDB)

	gdb = gdb.Debug()
	tx := gdb.Begin()

	var zimTom database.Pet
	result := tx.First(&zimTom, "nome = ?", "ZimTom")

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("1ª execução: registro não encontrado, foi inserido")
		tx.Create(database.Pets)
		tx.Commit()
	} else {
		fmt.Println(result)
		fmt.Println(zimTom.Nome, zimTom.Idade)
		tx.Rollback()
	}

}
