/* 4. Цикл for с break: Напишите цикл for, который выводит числа от 1 до 100, но
прерывает цикл, как только сумма выводимых чисел превысит 50. */

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {

		if i > 50 {
			break
		}
		fmt.Println(i)
	}
}
