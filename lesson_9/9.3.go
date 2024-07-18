/* 3. Изменение значений полей структуры: Создайте экземпляр структуры
из предыдущего задания, измените значение одного из полей и выведите
все поля на экран. */

package main

import (
	"fmt"
)

func main() {

	type Person struct {
		Name      string
		Age       int
		IsMarried bool
	}

	w := Person{
		Name:      "Иванов",
		Age:       33,
		IsMarried: true,
	}

	w.Age = 35

	fmt.Println(w, w.Name, w.Age, w.IsMarried)
}
