/* 4. Использование нескольких значений return: Создайте функцию, которая
принимает два числа и возвращает их сумму и разность как два отдельных значения. */

package main

import "fmt"

func mkv(a, b int) (int, int) {

	return (a + b), (a - b)
}

func main() {

	fmt.Println(mkv(9, 5))
}
