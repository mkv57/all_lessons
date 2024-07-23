/* 8. Конкатенация констант и переменных: Создайте переменную и
константу с типами string, присвойте им значения и выведите результат их
конкатенации.
*/

package main

import "fmt"

const a = "Hello "

var b = "world"

func main() {

	fmt.Println(a + b)

	c := a + b

	fmt.Println(c)
}
