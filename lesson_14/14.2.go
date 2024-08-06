/* 2. Десериализация JSON: Создайте JSON-строку, представляющую данные
о продукте. Десериализуйте эту строку в структуру Product и выведите
поля структуры на экран. */

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

	var deCarJison Product

	err = json.Unmarshal(carJison, &deCarJison)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	fmt.Println(deCarJison)

}
