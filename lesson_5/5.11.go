/* 11. Одинаковый элемент в слайсе: Создайте два слайса типа int со
значениями 1, 2, 6, 5, 8 и 9, 5, 3, 4, 1 соответственно. С помощью циклов
выведите значения, которые повторяются в слайсах. Попробуйте слайсы
с разными значениями и длиной для такого же итога. */

package main

import "fmt"

func main() {

	a := []int{1, 2, 6, 5, 8}

	b := []int{9, 5, 3, 4, 1}

	for i := 0; i < len(a); i++ {

		for j := 0; j < len(b); j++ {

			if a[i] == b[j] {

				fmt.Println(a[i]) // Всё работает с разными длинами слайсов
			}
		}
	}
}
