// Fibonacci algo from tour of go
package main

import "fmt"

func fib(n int, c chan int) {
	x, y := 0, 1

	// for loop until the buffered channel is full
	for i := 0; i < n; i++ {
		fmt.Println("x after i", i, "interations is:", x)
		c <- x        // x send to channel c
		x, y = y, x+y // changes the x and y for the fibonacci operation
	}
	close(c) // close channel x after for loop, so that the receiver knows that there will be no values to receive
}

func main() {
	c := make(chan int, 10)
	fmt.Println(cap(c)) // cap of c, which is 10

	go fib(cap(c), c)

	// for loop until channel c is closed
	for i := range c {
		fmt.Println(i)
	}

}
