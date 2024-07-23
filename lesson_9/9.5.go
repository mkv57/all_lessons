/* 5. Сравнение структур: Создайте два экземпляра одной и той же
структуры. Сравните их с помощью оператора == и выведите результат. */

package main

import (
	"fmt"
)

type www struct {
	Name string
	Age  int
}

func main() {

	r := www{
		"ok",
		45,
	}

	/*w := www{
		"ok",
		44,
	}*/
	w := www{
		"ok",
		45,
	}

	var s bool = r == w

	fmt.Println(s) // если зачения внутри полей совпадают, то правда, если нет, то ложь
}
