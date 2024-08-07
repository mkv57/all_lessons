package main

import (
	"fmt"
)

/*
	func Price(a, b, c int) int {
		return a * b * c
	}

	var size = []int{
		36,
		37,
	}

var a, b, c int
var n = Price(a, b, c)
*/
var models = []string{
	"model 1", "model 2",
}

var brend = []string{
	"Nike", "Adidas", "Puma",
}
var colors = []string{
	"red",
	"green",
}

func color(uu string) int {

	if uu == "red" {
		return 0
	} else if uu == "green" {
		return 1
	}
	return 5
}

func model(uu string) int {

	if uu == "model 1" {
		return 0
	} else if uu == "model 2" {
		return 1
	}
	return 5
}

func brend1(u string) int {

	if u == "Nike" {
		return 0
	} else if u == "Adidas" {
		return 1
	} else if u == "Puma" {
		return 2
	}
	return 5
}

var jBrend = brend1("Adidas")
var jModel = model("model 2")
var jColor = color("red")

func main() {

	fmt.Println(ID_00002)
}

var ID_00002 = map[string]string{

	"brend": brend[jBrend],
	"model": models[jModel],
	"color": colors[jColor],
}
