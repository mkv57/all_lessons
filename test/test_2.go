package main

import (
	"fmt"
	"strconv"
)

func main() {

	n := strconv.Itoa(color["green"]) + strconv.Itoa(brend["Adidas"]) + strconv.Itoa(model["model_3"])

	fmt.Println(n)
	c, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}

var color = map[string]int{
	"red":   11,
	"green": 12,
	"black": 13,
}
var brend = map[string]int{
	"Nike":   11,
	"Adidas": 12,
	"Puma":   13,
}
var model = map[string]int{
	"model_1": 11,
	"model_2": 12,
	"model_3": 13,
}
