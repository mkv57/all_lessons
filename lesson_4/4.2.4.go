/* 4. Создание слайса с помощью make: Создайте слайс с помощью функции
make, добавьте в него несколько элементов и выведите эти элементы на
экран. */

package main

import "fmt"

func main() {

	a := make([]string, 0, 3)

	a = append(a, "1", "2", "Hello", "4")

	fmt.Println(a)

}
