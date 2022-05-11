package main

import "fmt"

func explainPointers() {

	giveMeSomeSpace()
	fmt.Println("Lets learn about Go pointers")
	printLine()

	fmt.Println("Pointers are defined by *type, like *string type")
	fmt.Println("You can access a pointer value by using * in front of the variable name")
	fmt.Println("A pointer has nil as default value")
	fmt.Println("We can use & in front a variable to get its memory address (pointer)")

	var s *int
	var s2 = 42

	fmt.Println("My pointer is", s)

	s = &s2

	fmt.Println("Now, my pointer is", s)
	fmt.Println("My pointer value is ", *s)

	fmt.Println("Lets increment pointer's value and check both variables")

	*s++

	fmt.Println(s, *s, s2, &s2)

	printLine()
}
