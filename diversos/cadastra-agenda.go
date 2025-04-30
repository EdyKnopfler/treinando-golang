package main

import (
	"com.derso/testify/agenda"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("agenda.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&agenda.Agenda{})
	db.AutoMigrate(&agenda.Cliente{})
	db.AutoMigrate(&agenda.Apontamento{})

	db.Transaction(func(tx *gorm.DB) error {
		testeAgenda := agenda.Agenda{
			Nome: "Agenda Teste",
		}

		cliente := agenda.Cliente{
			Nome: "ZimTom",
		}

		tx.Create(&testeAgenda)
		tx.Create(&cliente)
		return nil
	})
}
