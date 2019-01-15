package main

import (
	"fmt"
)

// Gloabl Variables, can used by any function
// Scope of Variables: "Go is lexically scoped using blocks"
var abc = "Hello World!"

func main() {
	/*---Declaration of Variables---*/
	//var y int
	//y = 4
	//var x int = 2
	//	var z string = "Hello"
	//fmt.Printf("%v + %v", x, y)

	/* var a string = "Hello"
	var b string = "World"
	fmt.Println(a + b)
	a = "Bye"
	fmt.Printf("%v %v", a, b) */

	/* var c string
	var d string
	c = "Hello"
	d = "Hello"
	fmt.Println("FirstLine " + c)
	c = c + "World"
	fmt.Println("SecondLine " + c)
	d += "Univers"
	fmt.Println("ThirdLine" + d) */

	// Type Inference
	/* var e = "Hello"
	f := "World"
	g := 4.5
	fmt.Println("E: ", reflect.TypeOf(e))
	fmt.Println("F: ", reflect.TypeOf(f))
	fmt.Println("G: ", reflect.TypeOf(g))
	fmt.Println(e, f) */

	/* fmt.Print("Enter a number: ")
	var input float64
	// Read sdtin and %f for float variables
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output) */

	// Test example to convert fahrenheit in Celsius
	//fahrenheitCelsius()

	// Test example to convert feet in meters
	//feetMeter()

	// Example that swaps 2 variables
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Println(convert(5))
}

func fahrenheitCelsius() {
	fmt.Print("Enter a number: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	celsius := ((fahrenheit - 32) * 5 / 9)
	fmt.Println(celsius, "°C")
}

func feetMeter() {
	fmt.Print("Enter a length in feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meter := feet * 3.281
	fmt.Println(meter, "m")
}

// function that requires 2 args of type string and returns 2 strings (string, string)
func swap(x, y string) (string, string) {
	return y, x
}

// named required variable, returns x and y as int
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// "naked" return -> returns the named vars, in this case x and y
	return
}

func convert(before int) (after float64) {
	fmt.Printf("Type before: %T\n", before)
	after = float64(before)
	fmt.Printf("Type after: %T\n", after)
	return
}
