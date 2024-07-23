/* 7. Работа с bool (продвинутый уровень): Создайте переменные типа bool и
присвойте им разные значения. Напишите программу, которая выводит
результат логической операции "отрицание" (!) для каждой из
переменных. */

package main

import "fmt"

const w string = "Hello IT Academy"

func main() {

	fmt.Println(w)

	var a bool = false
	var b bool = true

	if a != true && b != true {

		fmt.Println("a and b  = false")
	} else if a == b {

		fmt.Println("a and b = true")

	} else if a != true {
		fmt.Println("a = false and b = true")
	} else {
		fmt.Println("a = true and b = false")
	}

}
