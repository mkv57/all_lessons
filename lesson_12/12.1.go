/* 1. Возвращение ошибки из функции: Напишите функцию, которая
возвращает ошибку. Вызовите эту функцию и обработайте возвращенную
ошибку. */

package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	m := []int{2, 3, 7}

	c := "56r7"

	i, err := strconv.Atoi(c)
	if err != nil {
		fmt.Errorf("kbhsb : %d", i)
		log.Fatal(err)
		return
	}

	m[0] = i

	fmt.Println(m)
}
