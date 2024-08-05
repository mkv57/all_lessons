package main

import "fmt"

var n int = 86 // курс доллара

var tesla Car = Car{500000, n}
var boing Aircraft = Aircraft{10000000, n}

type Price interface {
	move()
	//move1(x,y int) int
}

type Car struct {
	price int // цена в долларах
	rate  int // курс доллара
}
type Aircraft struct {
	price int // цена в долларах
	rate  int // курс доллара
}

func (c Car) move() {
	fmt.Print("Автомобиль стоит ", c.price*c.rate)
}
func (a Aircraft) move() {
	fmt.Print("Самолет стоит ", a.price*a.rate)
}

func price1(g Price) {
	g.move()
	fmt.Println(" рублей")
}

// использование интефейса помогает сократить кол-во функций

/*func priceCar(c Car) {
	c.move()
	fmt.Println(c.price*c.rate, " рублей")
}
func priceAircraft(a Aircraft) {
	a.move()
	fmt.Println(a.price*a.rate, " рублей")
}
*/
func main() {

	price1(tesla)
	price1(boing)
}
