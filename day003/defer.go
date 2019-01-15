package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	// Stack nr 0
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()
	fmt.Println("calling in g")
	g(0)
	// will never be executed, because the program paniced and was terminiated after recover
	fmt.Println("returned normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	// Stack nr 1+2+3+4
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
