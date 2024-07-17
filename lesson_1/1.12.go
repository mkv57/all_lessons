/* 12. Конкатенация: Создайте две переменные типа string, одну со значением
“Hello”, другую со значением “World!”. Выведите результат слияния двух
строк с помощью оператора “+”. */

package main

import "fmt"

const w string = "Hello IT Academy"

func main() {

	fmt.Println(w)
	a := "Hello"
	b := "World"
	c := a + b

	fmt.Println(c)
	fmt.Println(a, b)
	fmt.Println(a + b)

}
