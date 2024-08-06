/* 2. Обработчики маршрутов: Создайте несколько обработчиков маршрутов.
Например, /user , который будет возвращать информацию о пользователе,
и /product , который будет возвращать информацию о продукте.*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
type Product struct {
	Name string `json:"name"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	user := User{Name: "Konstantin", Age: 57, Email: "mkv2008@yandex.ru"}

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userJSON)
	if err != nil {
		log.Fatal(err)
	}

}

func userHandler1(w http.ResponseWriter, r *http.Request) {

	product := Product{Name: "Car"}

	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(productJSON)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/product", userHandler1)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
