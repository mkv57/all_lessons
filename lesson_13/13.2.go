/* 2. Передача указателя в функцию: Напишите функцию, которая принимает
указатель на int и инкрементирует значение, на которое указывает
указатель. */

package main

import (
	"fmt"
)

func main() {

	a := 27
	p := &a

	incrementIntegerByPointer(p)
	fmt.Println(a)
}

func incrementIntegerByPointer(a *int) {

	*a++
}
