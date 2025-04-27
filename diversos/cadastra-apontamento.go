package main

import (
	"time"

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

	idCliente := uint64(1)
	idAgenda := uint64(1)

	var cliente agenda.Cliente
	db.First(&cliente, idCliente)

	loc, _ := time.LoadLocation("America/Sao_Paulo")

	apontamento := agenda.Apontamento{
		Inicio:   time.Date(2025, time.May, 1, 13, 30, 0, 0, loc),
		Fim:      time.Date(2025, time.May, 1, 15, 30, 0, 0, loc),
		IdAgenda: idAgenda,
		Cliente:  cliente,
	}

	db.Create(&apontamento)
	db.Commit()
}
