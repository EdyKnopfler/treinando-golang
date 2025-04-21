package main

import (
	"fmt"
	"math"
)

func potencia(n, exp int) int {
	if exp == 0 {
		return 1
	}

	e := exp / 2
	r := potencia(n, e)

	if exp%2 == 0 {
		return r * r
	} else {
		return r * r * n
	}
}

func main() {
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if potencia(i, j) != int(math.Pow(float64(i), float64(j))) {
				panic(fmt.Sprintf("Faiô: %d^%d", i, j))
			}
		}
	}

	fmt.Println("Passô tudo")
}
