/* 8. Слайс как значение: Создайте map, где значением будут слайс типа int.
Создайте несколько пар ключ-значение в данной map. Выведите
результат.*/

package main

import "fmt"

func main() {

	a := map[string][]int{}

	a["qqq"] = []int{4}
	a["www"] = []int{5}

	fmt.Println(a)
}
