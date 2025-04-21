/*
PROMPT NO CHAT GP-TOLA

Dado um csv com as colunas: data-hora (em iso 8601), cliente, item comprado, quantidade, preço unitário.
Faça um programa em golang que lê esse arquivo e dê os totais vendidos:
(1) por data, sendo a data no timezone America/Sao_Paulo independentemente de como está no CSV (por exemplo, supondo que o CSV esteja em UTC, quero os totais por data local)
(2) por cliente
(3) por item
(4) por item, por data
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Sale struct {
	DateTime  time.Time
	Client    string
	Item      string
	Quantity  int
	UnitPrice float64
}

func parseCSV(filename string) ([]Sale, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var sales []Sale
	for _, row := range rows[1:] {
		dateTime, err := time.Parse(time.RFC3339, row[0])
		if err != nil {
			return nil, err
		}
		quantity, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, err
		}
		unitPrice, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			return nil, err
		}

		sales = append(sales, Sale{
			DateTime:  dateTime,
			Client:    row[1],
			Item:      row[2],
			Quantity:  quantity,
			UnitPrice: unitPrice,
		})
	}
	return sales, nil
}

func processSales(sales []Sale) {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	totalByDate := make(map[string]float64)
	totalByClient := make(map[string]float64)
	totalByItem := make(map[string]float64)
	totalByItemByDate := make(map[string]map[string]float64)

	for _, sale := range sales {
		localDate := sale.DateTime.In(location).Format("2006-01-02")
		total := float64(sale.Quantity) * sale.UnitPrice

		totalByDate[localDate] += total
		totalByClient[sale.Client] += total
		totalByItem[sale.Item] += total
		if totalByItemByDate[sale.Item] == nil {
			totalByItemByDate[sale.Item] = make(map[string]float64)
		}
		totalByItemByDate[sale.Item][localDate] += total
	}

	fmt.Println("Total por data:")
	for date, total := range totalByDate {
		fmt.Printf("%s: R$ %.2f\n", date, total)
	}

	fmt.Println("\nTotal por cliente:")
	for client, total := range totalByClient {
		fmt.Printf("%s: R$ %.2f\n", client, total)
	}

	fmt.Println("\nTotal por item:")
	for item, total := range totalByItem {
		fmt.Printf("%s: R$ %.2f\n", item, total)
	}

	fmt.Println("\nTotal por item e data:")
	for item, dates := range totalByItemByDate {
		fmt.Printf("%s:\n", item)
		for date, total := range dates {
			fmt.Printf("  %s: R$ %.2f\n", date, total)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <arquivo.csv>")
		return
	}

	sales, err := parseCSV(os.Args[1])

	if err != nil {
		fmt.Println("Erro ao processar CSV:", err)
		return
	}

	processSales(sales)
}
