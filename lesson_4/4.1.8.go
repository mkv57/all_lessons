/*
 8. Изменение нескольких элементов массивы: Создайте массив типа int с

элементами 1, 2, 3, 4, измените их порядок в массиве на 4, 3, 2, 1 с
помощью изменения элементов по индексу.
*/
package main

import "fmt"

func main() {

	a := [4]int{1, 2, 3, 4}
	fmt.Println(a)

	a[0], a[1], a[2], a[3] = a[3], a[2], a[1], a[0]

	fmt.Println(a)
}
