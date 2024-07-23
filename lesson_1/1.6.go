/* 6. Математические операции с float: Объявите две переменные типа
float64, присвойте им значения и выполните различные математические
операции. Обратите внимание на точность результатов */

package main

import "fmt"

const w string = "Hello IT Academy"

func main() {

	fmt.Println(w)

	var a float64 = 57.3
	var b float64 = 21
	var c float64 = a + b
	d := a * b
	e := a / b
	f := d + e
	fmt.Println("", a, "\n", b, "\n", c, "\n", d, "\n", e, "\n", f)
	// меняется количество знаков после запятой, но не меняется количество знаков
}
