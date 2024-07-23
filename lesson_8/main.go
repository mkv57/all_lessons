/* 1. Создание и использование собственного пакета: Создайте свой
собственный пакет с функцией, которая возвращает строку.
Импортируйте и используйте этот пакет в другой программе */

package main

import (
	"fmt"
	"lesson_8/test_1"
)

func main() {

	fmt.Println(test_1.Log("rrr"))
	test_1.Log("This is a message from test_1 www")
	fmt.Println(test_1.Log_1(4, 7))
}
