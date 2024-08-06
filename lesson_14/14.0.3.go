/* 3. Обработка ошибок: Добавьте обработку ошибок в ваш сервер. Например,
если обработчик не может декодировать JSON, он должен возвращать
соответствующий код ошибки и сообщение. Подумайте, какой код лучше
всего подходит под эту ситуацию. */

package main

import (
	"encoding/json"
	"fmt"
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
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err, "ошибка конвертации в json")
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userJSON)

	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/hello", userHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
