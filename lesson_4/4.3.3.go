/* 3. Получение значения по ключу из map: Создайте map, добавьте в нее
несколько пар ключ-значение. Затем получите и выведите на экран
значение по одному из ключей. */

package main

import "fmt"

func main() {

	a := map[string]int{}

	a["fff"] = 2
	a["ddd"] = 9

	fmt.Println(a["fff"])

	n := a["fff"] // если под "Затем получите" имеллось ввиду вытащить значение

	fmt.Println(n)
}
