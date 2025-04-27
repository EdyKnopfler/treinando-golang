package main

import (
	"fmt"

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

	var apontamentos []agenda.Apontamento
	db.Preload("Cliente").Find(&apontamentos, agenda.Apontamento{IdCliente: 1})

	for _, apontamento := range apontamentos {
		fmt.Println(apontamento.Cliente.Nome, apontamento.Inicio, apontamento.Fim, apontamento.Status)
	}

	var cliente agenda.Cliente
	db.Preload("Apontamentos").First(&cliente, 1)
	fmt.Println(cliente.Nome, cliente.Apontamentos)
}
