/* 5. Оператор выбора switch с несколькими значениями для case: Создайте
слайс из чисел. Напишите цикл, который использует оператор switch для
проверки каждого числа и выводит сообщение для четных и нечетных
чисел. */

package main

import (
	"fmt"
)

func main() {

	a := []int{2, 5, 44, 3, 0, 6, 98}

	for i := 0; i < len(a); i++ {

		/*switch a[i] % 2 {

		case 0:
			fmt.Println(a[i], " чётное")
		default:
			fmt.Println(a[i], " не чётное")	//  Что с "o" делать без "if", пока не понимаю
		}*/
		switch 0 {

		case a[i]:
			fmt.Println(a[i], " равно 0")
		case a[i] % 2:
			fmt.Println(a[i], " чётное")
		default:
			fmt.Println(a[i], " не чётное")
		} // xчто-то придумал

	}

}
