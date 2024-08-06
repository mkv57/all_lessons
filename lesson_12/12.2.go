/* 2. Создание новой ошибки: Используйте функцию errors.New для создания
новой ошибки. Верните эту ошибку из вашей функции и обработайте её. */

package main

import (
	"errors"
	"fmt"
)

var err = errors.New("ОШИБКА")

func www(x int) error {
	return err
}

func main() {

	r := www(34)

	if r != nil {
		fmt.Println(r)
	}
}
