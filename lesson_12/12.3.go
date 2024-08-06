/* 3. Проверка на наличие ошибки: Напишите функцию, которая возвращает
ошибку в некоторых случаях. Вызовите эту функцию и проверьте, есть ли
ошибка, прежде чем продолжать выполнение кода.*/

package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	с, err := myfunc()
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("Введено число:", с)
}

func myfunc() (int, error) {
	var a int
	var myError = errors.New("Введено не корректное число")
	fmt.Scan(&a)
	if a < 9 && a > 7 {
		return a, myError
	}
	return a, nil
}
