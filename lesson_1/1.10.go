/* 10. Комплексное задание (⭐ сложное): Создайте программу, которая
принимает на вход два значения разных типов (например, string и int),
преобразует одно из значений в другой тип и производит
математическую операцию. Результат должен выводиться на экран. Для
выполнения этого задания вам потребуется самостоятельно изучить
функцию Scan. */

package main

import "fmt"

const w string = "Hello IT Academy"

func main() {

	fmt.Println(w)

	var name string
	var age int
	fmt.Println("What is your name")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
	fmt.Println("How old are you ?")
	fmt.Scan(&age)
	fmt.Println("You are " + fmt.Sprint(age) + " years")

}
