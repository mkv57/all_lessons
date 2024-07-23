/* 1. Объявление массивов: Создайте массив из пяти целых чисел и
присвойте каждому элементу значение. Выведите каждый элемент на
экран отдельно.*/

package main

import "fmt"

func main() {

	s := [5]int{}

	s[0] = 2
	s[1] = 3
	s[2] = 5
	s[3] = 7
	s[4] = 11

	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	fmt.Println(s[3])
	fmt.Println(s[4])

}
