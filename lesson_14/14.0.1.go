/* 1. Создание простого HTTP сервера: Создайте HTTP сервер, который будет
слушать на порту 8080 и возвращать приветственное сообщение в
формате JSON при обращении по пути /hello .*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User1 struct {
	Greetings string `jsom:"greetings"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	hello := User1{Greetings: "Hello"}

	userJSON, err := json.Marshal(hello)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userJSON)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/hello", userHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
