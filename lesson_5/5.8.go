/* 8. Цикл for-range с строкой: Создайте строку. Напишите цикл for-range,
который выводит каждый символ строки на экран. */

package main

import (
	"fmt"
)

func main() {

	for _, v := range "wer23" {

		fmt.Printf("%c\n", v)

	}
}
