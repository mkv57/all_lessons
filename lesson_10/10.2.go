/* 2. Методы, изменяющие значения: Добавьте к структуре "Rectangle" метод
"DoubleSize", который удваивает размеры прямоугольника. */

package main

import (
	"fmt"
)

type Recktangle struct {
	Width  float64 // длинна 1ой стороны прямоугольника
	Height float64 // длинна 2ой стороны прямоугольника
}

func (p Recktangle) DoubleSize() Recktangle {

	return Recktangle{
		Width:  p.Width * 2,
		Height: p.Height * 2,
	}

}

func main() {

	p := Recktangle{Width: 35.2, Height: 34.2}

	fmt.Println(p.DoubleSize())
}
