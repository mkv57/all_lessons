/* 6. Сложная структура (⭐ сложное): Создайте структуру Order , которая
содержит поле Products , являющееся срезом структур Product .
Сериализуйте и десериализуйте экземпляр структуры Order .*/

package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Products []Product `json:"products"`
}

type Product struct {
	Name string `json:"name"`
}

func main() {

	Product_1 := make([]Product, 2)
	Product_1[0] = Product{Name: "tomato"}
	Product_1[1] = Product{Name: "cucumber"}

	order := Order{Products: Product_1}

	orderJson, err := json.Marshal(order)
	if err != nil {
		fmt.Print(err)
		return
	}

	var deconvOrder Order

	err = json.Unmarshal(orderJson, &deconvOrder)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(Product_1)
	fmt.Println(string(orderJson))
	fmt.Println(deconvOrder)

}
