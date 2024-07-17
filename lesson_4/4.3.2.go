/* 2. Изменение значений в map: Создайте map с несколькими парами ключ-значение. Измените одно из значений и выведите всю map на экран. */

package main

import "fmt"

func main() {

	a := map[string]int{"sss": 3, "ttt": 7, "ggg": 6}

	a["ttt"] = 1

	fmt.Println(a)
}
