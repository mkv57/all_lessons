/* 3. Цикл for с continue: Модифицируйте предыдущий цикл так, чтобы он
использовал команду continue для пропуска нечетных чисел. */

package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {

		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
