package main

import "fmt"
import "runtime"

func main() {

	// classic for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// for loop without init and post statements
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	// while loop
	/* for i < 16 {
		fmt.Println(i)
	} */

	// if statement
	if true {
		fmt.Println("Statement is true!")
	}

	if false {
		fmt.Println("Statement is false!")
	}

	// short statement before condition
	x := 10
	a := 5
	if v := x + a; v < 20 {
		fmt.Println("v is less than 20")
	}
	// short statements can be used in the else if block and variables declared in the if block are available in the else block
	if v := x + a; v < 10 {
		fmt.Println("v is less than 10")
	} else if v == 15 {
		fmt.Println("v is equal 15")
	}

	// switch case statement; example from Tour of Go
	fmt.Printf("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}

	// switch true == if else
	/* switch {
	case:
	case:
	default:
	} */

	// defer functions will be evaluated, but not executed until the surrounding function returns
	// LiFo Stack
	defer fmt.Println("Hello")
	fmt.Println("\nWorld")

	deferTest()

}

func deferTest() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
