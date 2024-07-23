/* 6. Структуры и функции (⭐ сложное): Создайте функцию, которая
принимает два аргумента типа "Book" и возвращает книгу с большим
количеством страниц. Протестируйте эту функцию с несколькими
экземплярами структуры "Book" */

package main

import (
	"fmt"
)

type Book struct {
	Pages int
}

func Www(Tom_1, Tom_2 Book) Book {

	if Tom_1.Pages > Tom_2.Pages {

		return Tom_1
	}

	return Tom_2

}

func main() {

	Tom_1 := Book{
		123,
	}
	Tom_2 := Book{
		345,
	}

	fmt.Println(Www(Tom_1, Tom_2))
}
