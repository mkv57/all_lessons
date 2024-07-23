/* 4. Конвертация типов: Создайте переменную типа string, содержащую
число. Преобразуйте эту строку в int и выполните некоторые
математические операции */

package main

import (
	"fmt"
	"strconv"
)

const w string = "Hello IT Academy"

func main() {

	fmt.Println(w)

	var k string = "57"
	var a, err = strconv.Atoi(k)
	var b int = 21
	var c int = a + b
	var d int = a - b
	var e int = a * b
	var f int = a / b
	if err != nil {

		fmt.Println("error")

	} else {
		fmt.Println("", a, "\n", b, "\n", c, "\n", d, "\n", e, "\n", f)
	}

}
