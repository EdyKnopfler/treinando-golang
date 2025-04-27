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

	testeAgenda := agenda.Agenda{
		Nome: "Agenda Teste",
	}

	cliente := agenda.Cliente{
		Nome: "ZimTom",
	}

	db.Create(&testeAgenda)
	db.Create(&cliente)

	db.Commit()
}
