/*
 5. Многомерные массивы: Создайте двумерный массив (массив массивов)

и присвойте каждому элементу значение. Выведите все значения на
экран.
*/
package main

import "fmt"

func main() {

	a := [3][3]int{} // создан двумерный массив из 3 строк и 3 столбцов

	a[0][0] = 1 // присвоено значение элементу
	a[0][1] = 2
	a[0][2] = 3
	a[1][0] = 4
	a[1][1] = 4
	a[1][2] = 5
	a[2][0] = 6
	a[2][1] = 7
	a[2][2] = 8

	fmt.Println(a)

	b := [2][3]int{{1, 2, 3}, {4, 5, 6}} // создан двумерный массив из 2 строк и 3 столбцов с назначенными зачениями каждому элементу

	fmt.Println(b)

}
