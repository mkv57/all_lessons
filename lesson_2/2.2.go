/*
 2. Объявление локальных констант: Внутри функции объявите локальную

константу и используйте ее. Выведите значение этой константы.
*/
package main

import "fmt"

func main() {

	const a, b = 57, "Hello" // Локальные константы.

	fmt.Println(a, b)
}
