/* 3. Использование тегов: Создайте структуру Employee с полями ID , Name , и
Product . Используйте теги для изменения имен полей в JSON.
Сериализуйте экземпляр структуры в JSON и убедитесь, что имена полей
соответствуют тегам.*/

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	user := Employee{ID: 333, Name: "Ivanov"}

	userJison, err := json.Marshal(user)

	if err != nil {
		fmt.Println("Ошибка сериализации", err)
		return
	}

	fmt.Println(string(userJison))
}
