package main

import "fmt"

// takes channel as argument
func sum(s []int, c chan int) {
	sum := 0
	// adds the numbers in the slice
	for _, v := range s {
		sum += v
	}
	fmt.Println("Computation of", s, "must be:", sum)
	// sends sum to channel c
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// create a channel; must be created before you can receive or send something
	// unbufferd channel
	c := make(chan int)
	// non-blocking goroutines
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// receives from channel c
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

/*
	ch := make(chan int)

	// receives from channel ch
	// and assign value to sum
	sum := <- ch

	// sends to channel ch
	ch <- ch
*/

// unbuffered channel
/*
	c := make(chan int)

	-> always full channel
	-> blocks until a receiver receives and a sender sends a value
	Writes to a unbuffered channel happens only, when there is a routine to read from the channel, else write operation is blocked
	forever and leads to deadlock
*/

// buffered channel
/*
	c := make(chan int, 2)

	buffered channel can receive n send operations until it makes an write operation
	same as unbuffered channel, but the Sender is not blocked until the buffer is full, receivers is not necessarily synchronized with
	every write operation
*/

// closing a channel
/*
	with close(channel) you can close a channel.
	Best if the sender is closeing a channel. Sending to a closed channel causes panic

	var, ok := <-ch
	ok is a second parameter to check if a channel is closed or not
	if ok is false, if no more values to receive and the channel is closed

	for i := range ch
	-> receives values from the channel repeatedly until it's closed

*/
