/* 1. Создание и использование указателей: Создайте переменную и
указатель на эту переменную. Измените значение переменной, используя
указатель. */

package main

import (
	"fmt"
)

func main() {

	var n int = 4
	p := &n // создал указатель для переменной с адресом переменной n
	*p = 5  // присвоил новое значение для переменной n

	fmt.Println(n, p)
}
