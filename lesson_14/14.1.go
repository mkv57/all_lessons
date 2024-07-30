/* 1. Сериализация структуры: Создайте структуру Product с полями Name ,
Price и Quantity . Создайте экземпляр этой структуры и сериализуйте его в
JSON.*/

package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func main() {

	car := Product{Name: "BMW", Price: 10000000, Quantity: 4}

	carJison, err := json.Marshal(car)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(car, string(carJison))
}
