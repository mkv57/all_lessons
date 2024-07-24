/* 4. Использование fmt.Errorf для создания ошибки: Используйте функцию
fmt.Errorf для создания новой ошибки с форматированным сообщением.
Верните эту ошибку из вашей функции и обработайте её.*/

package main

import (
	"fmt"
	"log"
)

func main() {

	a, err := ww(24, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)

}
func ww(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("число: %d не должно быть  равно 0", y)
	}
	return x / y, nil
}
