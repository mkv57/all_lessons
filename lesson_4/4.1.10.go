/*
 10. Сохранение элемента массивы в переменную: Создайте массив типа int

с 3 элементами, также создайте 3 переменные разных типов (int, string,
float64). Сохраните элементы массива в переменные с помощью
конвертации. Выведите результат.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	m := [3]int{}

	a := 2
	b := 3.5
	c := "56r7"

	m[0] = a

	m[1] = int(b)

	i, err := strconv.Atoi(c)
	if err != nil {
		return
	}

	m[2] = i

	fmt.Println(m)
}
