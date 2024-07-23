/* 9. Изменение вместимости слайса: Создайте слайс произвольного типа с 6
элементами, узнайте его вместимость с помощью функции cap, опытным
путём определите в каком случае вместимость слайса будет 12, 24 и
более. */

package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(cap(a), a) // cap = 6

	a = append(a, 0)
	fmt.Println(cap(a), a) // cap = 12

	a = append(a, a...)
	fmt.Println(cap(a), a) // cap = 24

	a = append(a, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2)
	fmt.Println(cap(a), a) // cap > 24
}
