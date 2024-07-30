/* 7. Вложенные структуры (⭐ сложное): Создайте структуру Company , которая
содержит поле Employees , являющееся срезом структур Employee .
Сериализуйте и десериализуйте экземпляр структуры Company .*/

package main

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	Employees []Employee `json:"employees"`
}

type Employee struct {
	Name string `json:"name"`
}

func main() {

	Product_1 := make([]Employee, 2)
	Product_1[0] = Employee{Name: "tomato"}
	Product_1[1] = Employee{Name: "cucumber"}

	order := Company{Employees: Product_1}

	orderJson, err := json.Marshal(order)
	if err != nil {
		fmt.Print(err)
		return
	}

	var deconvOrder Company

	err = json.Unmarshal(orderJson, &deconvOrder)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(Product_1)
	fmt.Println(string(orderJson))
	fmt.Println(deconvOrder)

}
