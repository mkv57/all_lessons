/* 3. Обработка ошибок: Добавьте обработку ошибок в ваш сервер. Например,
если обработчик не может декодировать JSON, он должен возвращать
соответствующий код ошибки и сообщение. Подумайте, какой код лучше
всего подходит под эту ситуацию. */

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
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	_, err = w.Write(userJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")     // как это работает?
	//json.NewEncoder(w).Encode(hello)						// как это работает?

}

func main() {
	http.HandleFunc("/hello", userHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
