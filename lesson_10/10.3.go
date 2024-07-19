/* 3. Использование нескольких методов: Используйте методы "Area" и
"IsSquare" на нескольких экземплярах структуры "Rectangle". */

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

func (p Recktangle) Area() float64 {
	return p.Width * p.Height
}

func main() {

	p := Recktangle{Width: 34, Height: 34}
	c := Recktangle{Width: 55, Height: 44}
	fmt.Println(p.IsSquare(), c.IsSquare())
	fmt.Println(p.Area(), c.Area())
}
