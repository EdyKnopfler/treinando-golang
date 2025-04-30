package main

import (
	"fmt"
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
	loc, _ := time.LoadLocation("America/Sao_Paulo")

	db.Transaction(func(tx *gorm.DB) error {
		var cliente agenda.Cliente
		// Para salvar tudo junto e ficar consistente, pré-ccarregamos
		// Desnecessauro se o objetivo fosse somente inserir um apontamento sem ver os outros
		tx.Preload("Apontamentos").First(&cliente, idCliente)
		fmt.Println("Antes", cliente.Apontamentos)

		cliente.Apontamentos = append(cliente.Apontamentos, agenda.Apontamento{
			Inicio:   time.Date(2025, time.June, 1, 13, 30, 0, 0, loc),
			Fim:      time.Date(2025, time.June, 1, 15, 30, 0, 0, loc),
			IdAgenda: idAgenda,
		})

		fmt.Println("Depois", cliente.Apontamentos)
		tx.Save(&cliente)
		return nil
	})
}
