/* 1. Создание и использование структур: Создайте структуру,
представляющую собой "Book" с полями "Title", "Author" и "Pages".
Создайте экземпляр этой структуры, заполните поля и выведите их на
экран. */

package main

import "fmt"

func main() {

	type Book struct {
		Title  string
		Author string
		Pages  int
	}
	/*r := new(Book)
	r.Title = "инструкция"
	r.Author = "Иванов"
	r.Pages = 35 */

	r := Book{
		Title:  "Инструкция",
		Author: "Иванов",
		Pages:  35,
	}

	fmt.Println(r, r.Title, r.Author, r.Pages)

}
