/* 3. Использование bool: Объявите две переменные типа bool и присвойте им
значения. Напишите программу, которая выводит результат логической
операции "И" (&&) и "ИЛИ" (||) этих переменных.*/

package main

import "fmt"

const w string = "Hello IT Academy"

func main() {

	fmt.Println(w)

	var a bool = true
	var b bool = false

	var z1 bool = b && a
	var z2 bool = a || b
	var z3 bool = !a

	fmt.Println("", "z1 = ", z1, "\n", "z2 = ", z2, "\n", "z3 = ", z3, "\n")

}
