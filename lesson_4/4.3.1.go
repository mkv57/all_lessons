/* 1. Создание map: Создайте map, где ключом является строка, а значением -
целое число. Добавьте в нее несколько пар ключ-значение и выведите их
на экран. */

package main

import "fmt"

func main() {

	a := map[string]int{}

	a["fff"] = 2
	a["ddd"] = 9

	fmt.Println(a)
}
