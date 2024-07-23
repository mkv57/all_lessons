/* 10. Цикл в цикле: Используя два цикла выведите таблицу умножения. */

package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {

		for a := 1; a < 10; a++ {

			fmt.Print(i*a, "\t")

		}
		fmt.Println()
	}
}
