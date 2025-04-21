package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	texto := "世界"
	fmt.Println("Hello, ", texto)
	fmt.Println(len(texto), utf8.RuneCountInString(texto))
}
