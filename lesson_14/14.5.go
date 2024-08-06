/* 5. Обработка ошибок десериализации: Напишите функцию, которая
принимает JSON-строку и возвращает структуру. Обработайте
возможные ошибки при десериализации и выведите их на экран. */

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Produckt struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Conv(prod Produckt) []byte {

	carJson, err := json.Marshal(prod)
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
	return carJson
}
func DeConv(convEmployee []byte) Produckt {

	var deconvEmployee Produckt
	err := json.Unmarshal(convEmployee, &deconvEmployee)
	if err != nil {
		fmt.Print(err)
		log.Fatal()
	}
	return deconvEmployee
}

func main() {

	employee := Produckt{Name: "Ivanov", Age: 33}

	convEmployee := Conv(employee)

	deConvEmployee := DeConv(convEmployee)

	fmt.Println(deConvEmployee)
}
