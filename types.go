package main

import (
	"fmt"
	"math"
	"reflect"
)

func explainTypes() {

	giveMeSomeSpace()
	fmt.Println("Lets learn about Go types")
	printLine()

	fmt.Println("We have unasined int types")

	var a uint8 = 1
	var b uint16 = 2
	var c uint16 = 3
	var d uint32 = 4

	fmt.Println(a, "is", reflect.TypeOf(a), "also called byte")
	fmt.Println(b, "is", reflect.TypeOf(b))
	fmt.Println(c, "is", reflect.TypeOf(c))
	fmt.Println(d, "is", reflect.TypeOf(d))

	printLine()

	fmt.Println("We have sined int types")

	var e int8 = -5
	var f int16 = 6
	var g int32 = -7
	var h int64 = 8
	var i = math.MaxInt

	fmt.Println(e, "is", reflect.TypeOf(e))
	fmt.Println(f, "is", reflect.TypeOf(f))
	fmt.Println(g, "is", reflect.TypeOf(g))
	fmt.Println(h, "is", reflect.TypeOf(h))
	fmt.Println(i, "is the max int value, its value is", reflect.TypeOf(i))

	printLine()

	fmt.Println("We have float types")

	var j float32 = float32(PI)
	var k float64 = PI

	fmt.Println(j, "is", reflect.TypeOf(j))
	fmt.Println(k, "is", reflect.TypeOf(k))

	printLine()

	fmt.Println("We have boolean types")

	var l bool = false

	fmt.Println(l, "is", reflect.TypeOf(l))

	printLine()

	fmt.Println("And we have string types, with option to multiple line string")

	var m string = "Go"
	var n = `
	Hello
	Go
	`
	fmt.Println(m, "is", reflect.TypeOf(m))
	fmt.Println(n, "is", reflect.TypeOf(n))

	printLine()

	fmt.Println("But be careful, there is no char types.")

	var o = 'x'

	fmt.Println("So 'x' actually is", o, "of type", reflect.TypeOf(o))
	fmt.Println("And it's an int that maps the unicode table")

	printLine()
	var empty1 int
	var empty2 float64
	var empty3 bool
	var empty4 string
	var empty5 *int

	fmt.Println("For variable with value not defined we have the following defaulr value:")
	fmt.Println(reflect.TypeOf(empty1), "is", empty1)
	fmt.Println(reflect.TypeOf(empty2), "is", empty2)
	fmt.Println(reflect.TypeOf(empty3), "is", empty3)
	fmt.Print(reflect.TypeOf(empty4), " is ")
	fmt.Printf("%q", empty4)
	fmt.Println()
	fmt.Println(reflect.TypeOf(empty5), "is", empty5)

	printLine()
}
