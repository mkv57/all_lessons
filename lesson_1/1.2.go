/* 2. Математические операции: Создайте две переменные типа int и
присвойте им значения. Напишите программу, которая выводит результат
сложения, вычитания, умножения и деления этих переменных */

package main

import "fmt"

const w string = "Hello IT Academy"

func main() {

	fmt.Println(w)

	var a int = 57
	var b int = 21
	var c int = a + b
	var d int = a - b
	e := a * b
	f := a / b

	fmt.Println("", a, "\n", b, "\n", c, "\n", d, "\n", e, "\n", f)

}
