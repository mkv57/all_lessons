/* 4. Массивы с разными типами данных: Создайте три массива: один с
целыми числами, один с числами с плавающей точкой и один со
строками. Выведите все значения всех массивов. */

package main

import "fmt"

func main() {

	a := [3]int{1, 2, 2}
	b := [3]float64{1.7, 2.2, 2.5}
	c := [3]string{"hello", "world", "!!!"}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
