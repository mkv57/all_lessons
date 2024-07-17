package test_1

import "fmt"

func Log(mmm string) string {
	fmt.Println(mmm) // это лишнее)))
	return "message" // возращает строку
}

// Это тоже лишнее)))
func Log_1(x, y int) int { // возращает значение
	return x + y // возращает значение
}
