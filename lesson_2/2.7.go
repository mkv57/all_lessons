/* 7. Математические операции с константами и переменными: Создайте
переменную и константу с типами int, присвойте им значения и выполните
различные математические операции. Проведите аналогичные операции
для типа float64.
*/

package main

import "fmt"

const a int = 57

var b int = 21

const e float64 = 57.5

var f float64 = 21.4

func main() {

	c := a + b
	cc := a - b
	ccc := a * b
	cccc := a / b

	d := e + f
	dd := e - f
	ddd := e * f
	dddd := e / f

	fmt.Println(c)
	fmt.Println(cc)
	fmt.Println(ccc)
	fmt.Println(cccc, "\n")

	fmt.Println(d)
	fmt.Println(dd)
	fmt.Println(ddd)
	fmt.Println(dddd)
}
