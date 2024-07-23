/* 6. Вызов функции из другой функции: Создайте две функции и вызовите
одну из них из другой. */

package main

import "fmt"

func mkv(a int) {
	fmt.Println(a)
}

func msk() {
	mkv(3)
}
func main() {

	msk()

}
