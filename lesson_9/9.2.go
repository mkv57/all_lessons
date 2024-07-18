/* 2. Создание структур с разными типами данных: Создайте структуру
"Person" с полями "Name" (string), "Age" (int) и "IsMarried" (bool). Создайте
экземпляр этой структуры, заполните поля и выведите их на экран. */

package main

import (
	"fmt"
)

func main() {

	type Person struct {
		Name      string
		Age       int
		IsMarried bool
	}

	w := Person{
		Name:      "Иванов",
		Age:       33,
		IsMarried: true,
	}

	fmt.Println(w, w.Name, w.Age, w.IsMarried) // как я понял, то структуры не хранят значения, а хранят ячейки(поля) с типом данных,
	// заданных нами, и при вызове показывают вложенные нами значения, а где не вложены, то показывают пробел, если ячейка стринговая и 0, если int
}
