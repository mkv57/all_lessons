/*1. Объявление переменных: Создайте переменные различных типов (int,
float64, string, bool) и присвойте им значения. Выведите на экран значения
этих переменных.*/

package main

import "fmt"

const w string = "Hello IT Academy"

func main() {

	fmt.Println(w)

	var a int = 57
	var b float64 = 21
	var c string = "Have a good day"
	var d bool = true

	fmt.Println(c, "\n", a, "\n", b, "\n", d)

}
