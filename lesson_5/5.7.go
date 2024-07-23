/* 7. Цикл for-range с слайсом: Создайте слайс с несколькими элементами.
Напишите цикл for-range, который выводит каждый элемент слайса на
экран. */

package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4}

	for _, c := range a {

		fmt.Println(c)
	}
}
