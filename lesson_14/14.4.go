/* 4. Обработка ошибок сериализации: Напишите функцию, которая
принимает структуру и возвращает JSON-строку. Обработайте
возможные ошибки при сериализации и выведите их на экран.*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func convJson(Car Product) []byte {
	carJison, err := json.Marshal(Car)

	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
	return carJison
}

func main() {

	car := Product{Name: "BMW", Price: 10000000, Quantity: 4}

	fmt.Println(string(convJson(car)))
}
