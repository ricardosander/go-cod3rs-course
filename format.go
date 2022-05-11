package main

import (
	"fmt"
	"math"
)

func explainPrintAndFormating() {

	giveMeSomeSpace()
	fmt.Println("Lets learn about Go print and formating options")
	printLine()

	var radius = 3.2
	area := PI * math.Pow(radius, 2)

	sArea := fmt.Sprint(area)

	fmt.Printf("I can use Prinf to format a float like: PI is %.2f", PI)
	fmt.Println("\nOriginal PI value is", PI)
	fmt.Println("Or just Print a float like: radius is", radius)
	fmt.Println("But I also could convert a float to a string and then concat it to my text, like:")
	fmt.Println("Area is " + sArea)

	var d = 42
	var s = "Dont Pacnic"
	var b = true

	fmt.Println()

	fmt.Println("And I have a lot of formatter options:")
	fmt.Printf("\nDigit %d", d)
	fmt.Printf("\nFloat %f", area)
	fmt.Printf("\nFloat with 2 decimals %.1f", area)
	fmt.Printf("\nBoolean %t", b)
	fmt.Printf("\nString %s", s)
	fmt.Printf("\nAnything %v", radius)
	printLine()
}
