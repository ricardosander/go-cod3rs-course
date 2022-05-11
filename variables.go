package main

import (
	"fmt"
	"math"
)

func explainVariables() {

	var pi float64 = PI
	var radius = 3.2
	area := PI * math.Pow(radius, 2)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	var e, f bool = true, false

	g, h, i := 2, false, "oops"

	giveMeSomeSpace()
	fmt.Println("Lets learn about Go variables")
	printLine()

	fmt.Println("Please, check the code to see how to create variables")
	fmt.Println("But I will have to print all created constants and variables here ")
	fmt.Println("because Go doesn't compile when you define things you dont use: ")
	fmt.Println(a, b, c, d, e, f, g, h, i, area, pi)

	printLine()
}
