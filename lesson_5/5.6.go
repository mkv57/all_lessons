/* 6. Цикл for-range с map: Создайте map с несколькими парами ключ-
значение. Напишите цикл for-range, который выводит каждую пару ключ-
значение на экран. */

package main

import "fmt"

func main() {

	a := map[string]int{"qqq": 11, "www": 22, "eee": 33}

	for b, c := range a {

		fmt.Println(b, c)
	}
}
