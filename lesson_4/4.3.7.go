/* 7. Массив как ключ: Создайте два массива одного типа с одинаковой
длиной. Создайте map, в которой ключом являются созданные массивы,
тип значений задайте произвольно, попробуйте наполнить созданную map
разными данными.
Попробуйте использовать слайс как ключ в мапе, объясните итог. */

package main

import (
	"fmt"
)

func main() {

	// Вариан №1 (слайсы равны a = b)

	a := [3]int{}
	b := [3]int{} //так как "значения" массивов одинаковые, то и клучи, образованные от них будут одинаковыми
	//a := [3]int{1, 2, 3}
	//b := [3]int{1, 2, 3}   так как "значения" массивов одинаковые, то и клучи, образованные от них будут одинаковыми

	//c := map[[3]int]int{a: 1, b: 2}
	c := map[[3]int]int{}

	fmt.Println(c)

	c[a] = 6
	c[b] = 5

	fmt.Println(c)
	fmt.Println("") // разделил два блока для удобства, надеюсь так можно)))

	// Вариант №2 (слайсы не равны a != b)

	d := [3]int{0, 1, 2}
	e := [3]int{3, 4, 5}

	w := map[[3]int]int{d: 8, e: 9}

	fmt.Println(w)

	w[a] = 6
	w[b] = 7

	fmt.Println(w)

	/* f := map[[]]int{}  длина слайса может меняться, а она является частью типа слайса,
	соответственно и тип, при изменении длины, будет меняться.
	Наверно по этому тип ключа не может быть задан в таком виде.
	*/
}
