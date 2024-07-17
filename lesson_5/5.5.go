/* 5. Цикл for-range с массивом: Создайте массив с несколькими элементами.
Напишите цикл for-range, который выводит каждый элемент массива на
экран. */

package main

import "fmt"

func main() {

	a := [4]int{1, 2, 3, 4}

	for _, b := range a { //

		fmt.Println(b)
	}
}
