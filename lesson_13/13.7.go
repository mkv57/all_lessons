/* 7. Указатели и nil (⭐ сложное): Исследуйте, что произойдет, если
попытаться создать переменную с типом *int а затем разыменовать ее.
Придумайте вариант, который безопасно обрабатывает такую ситуацию. */

package main

import (
	"fmt"
	"error"
)

func main() {

	a := new(int)
	c := &a
	n := 45
	*c = &n
	fmt.Println(a)(n int, err )
	//fmt.Println(a, c, n, *c, *a)
	if err = !nil {
		return error
	}

	*a = 88

	fmt.Println(a, c, n, *c, *a)
}
