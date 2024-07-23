/* 2. Цикл for с условием: Напишите цикл for, который выводит на экран все
четные числа от 1 до 20. */

package main

import "fmt"

func main() {

	// Вариант 1

	/*	for i := 2; i <= 20; i = i + 2 {

			fmt.Println(i)
		}

	} */

	//Вариант 2

	for i := 1; i <= 20; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
