/* 8. Конвертация типов с float: Используя преобразование типов,
преобразуйте float64 в int и наоборот. Обратите внимание на то, что
происходит с десятичной частью числа при преобразовании. */

package main

import "fmt"

const w string = "Hello IT Academy"

func main() {

	fmt.Println(w)

	var a float64 = 57.76
	var b int = int(a)
	var c float64 = float64(b) // потеря десятичной части

	fmt.Println("\n", a, "\n", b, "\n", c)

}
