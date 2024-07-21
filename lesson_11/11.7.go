/* 7. Оператор выбора switch с fallthrough (⭐ сложное): Создайте слайс из
чисел от 1 до 10. Напишите цикл, который использует оператор switch с
fallthrough для проверки каждого числа и выводит сообщение для
каждого случая. */

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(a); i++ {

		switch a[i] {
		case 1:
			fmt.Println("a[", i, "] равно 1")
		case 2:
			fmt.Println("a[", i, "] равно 2")
		case 3:
			fmt.Println("a[", i, "] равно 3", " включение подачи инструмента + с подачей сож")
			fallthrough
		case 4:
			fmt.Println("a[", i, "] равно 3 или 4", "подача сож")
		case 5:
			fmt.Println("a[", i, "] равно 5")
		case 6:
			fmt.Println("a[", i, "] равно 6")
		case 7:
			fmt.Println("a[", i, "] равно 7")
		case 8:
			fmt.Println("a[", i, "] равно 8")
		case 9:
			fmt.Println("a[", i, "] равно 9")
		case 10:
			fmt.Println("a[", i, "] равно 10")

		}

	}

	// Понимаю, что это, что-то не то, но вроде бы, условия задачи выполнены)))

}
