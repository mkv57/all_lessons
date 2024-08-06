/* 3. Возвращение указателя из функции: Напишите функцию, которая
возвращает указатель на новую переменную int. */

package main

import (
	"fmt"
)

func main() {

	fmt.Println(www(23))
}

func www(n int) *int {
	p := &n
	fmt.Println(n)
	return p
}
