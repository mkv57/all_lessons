/* 4. Создание структуры с вложенными структурами (⭐ сложное):
Создайте структуру "Person" с полем "Name" и вложенной структурой
"Address" с полями "Street", "City" и "Country". Добавьте метод
"FullAddress", который возвращает полный адрес как строку. */

package main

import (
	"fmt"
)

type Address struct {
	Street  string
	City    string
	Country string
}

type Person struct {
	Name Address
}

func (p Address) FullAddress() Address {

	return p
}

func main() {
	/*p := Person{ 				//этот вариант  не работает, когда запускаем "fmt.Println(p.FullAddress())" Пока не могу понять почему?
		Name: Address{
			Street:  "К.Маркса",
			City:    "Тихвин",
			Country: "Россия",
		},
	}*/
	p := Address{
		Street:  "К.Маркса",
		City:    "Тихвин",
		Country: "Россия",
	}

	fmt.Println(p.FullAddress())
}
