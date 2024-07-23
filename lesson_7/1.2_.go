/* 2. Использование нескольких пакетов: Импортируйте пакеты "math" и
"time", используйте "math" для вычисления квадратного корня текущего
времени в секундах.*/

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var t_Sec float64 = float64(time.Now().Second())          // кол-во секунд на данный момент времени
	var t_Min float64 = float64(time.Now().Minute() * 60)     // кол-во минут на данный момент времени
	var t_Hour float64 = float64(time.Now().Hour() * 60 * 60) // кол-во часов на данный момент времени в формате от 0 до 23 часов

	var t float64 = float64((time.Now().Second() + (time.Now().Minute())*60 + (time.Now().Hour())*60*60)) // меньше переменных, экономия оперативной памяти

	fmt.Println(math.Sqrt(t_Sec + t_Min + t_Hour))
	fmt.Println(math.Sqrt(t))

}
