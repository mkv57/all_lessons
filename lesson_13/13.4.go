/* 4. Указатели на структуры: Создайте структуру и указатель на эту
структуру. Измените значение поля структуры, используя указатель в
методе. */

package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) Change() {
	r.Width = 11
}

func main() {

	s := &Rectangle{
		Width:  22,
		Height: 33,
	}
	s.Change()

	fmt.Println(s, s.Height, s.Width)
}
