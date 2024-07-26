/* 6. Создание кастомного типа ошибки (⭐ сложное): Создайте свой
собственный тип ошибки, реализовав интерфейс error. Создайте функцию,
которая возвращает вашу кастомную ошибку, и обработайте её. */

package main

import (
	"fmt"
)

var r = MyError{"не правильно введено число"}

func main() {

	r := MyError{"не правильно введено число"}
	err := Error()
	if err != nil {
		fmt.Println(r)

	}

}

type MyError struct {
	Msg string
}

func Error() error {
	var c int

	_, err := fmt.Scan(&c)

	return err

}
