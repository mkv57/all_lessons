/* 4. Удаление значения из map: Создайте map, добавьте в нее несколько пар
ключ-значение. Затем удалите одну из пар с помощью функции delete и
выведите всю map на экран. */

package main

import "fmt"

func main() {

	a := map[int]string{}

	a[1] = "qqq"
	a[2] = "www"
	a[3] = "eee"

	delete(a, 2)

	fmt.Println(a)
}
