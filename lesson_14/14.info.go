package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {

	user := User{Name: "Alice", Email: "alice@example.com", Age: 25} // Пример данных

	fmt.Println("Структура User:", user)

	userJSON, err := json.Marshal(user) // Сериализация структуры в JSON

	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}

	fmt.Println("Сериализованный JSON:", string(userJSON))

	var decodedUser User

	err = json.Unmarshal(userJSON, &decodedUser) // Десериализация JSON обратно в структуру
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)

		return
	}

	fmt.Println("Десериализованная структура:", decodedUser)
}
