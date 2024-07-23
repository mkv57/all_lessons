/* 1. Методы с разными возвращаемыми значениями: Добавьте к структуре
"Rectangle" метод "IsSquare", который возвращает bool, указывающий,
является ли прямоугольник квадратом. */

package main

import "fmt"

type Recktangle struct {
	Width  float64 // длинна 1ой стороны прямоугольника
	Height float64 // длинна 2ой стороны прямоугольника
}

func (p Recktangle) IsSquare() bool {

	var n bool = p.Width == p.Height //(1ый вариант) сравнение длинн сторон, если они равны, то прямоугольник является квадратом
	return n

	/*if p.Width == p.Height { // (2ый вариант) сравнение длинн сторон, если они равны, то прямоугольник является квадратом
		return true
	}
	return false */
}

func main() {

	p := Recktangle{Width: 35, Height: 34}
	fmt.Println(p.IsSquare())
}
