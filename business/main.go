package main

import (
	"io"
	"net/http"

	"com.derso/testify/business"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	mapData := business.BusinessRule(string(body))
	business.PrintMap(mapData, "")
}
