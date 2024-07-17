/* 3. Использование псевдонимов для пакетов: Импортируйте пакет
"math/rand" с псевдонимом и используйте его для генерации случайного
числа. */

package main

import (
	"fmt"

	koreny_2 "math/rand"
)

func main() {

	fmt.Println(koreny_2.Int())

}
