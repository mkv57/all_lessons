/* 4. Создание вложенных структур: Создайте структуру "Employee" с полем
"Name" и вложенной структурой "Job" с полями "Title" и "Salary".
Создайте экземпляр этой структуры, заполните все поля и выведите их на
экран. */

package main

import (
	"fmt"
)

type Job struct {
	Title  string
	Salary int
}
type Employee struct {
	Name Job
}

func main() {

	w := new(Job) // 1ый вариант
	w.Title = "ingenier"
	w.Salary = 300000

	r := new(Employee) // 2ой вариант
	r.Name = Job{
		Title:  "инженер",
		Salary: 250000, // А по какому варианту правильно заполнять поля?
	}

	fmt.Println(w, r)
}
