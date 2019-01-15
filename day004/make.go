package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlices("a", a)

	// cap of 5, but you can't see it, cause len is 0
	b := make([]int, 0, 5)
	printSlices("b", b)

	c := b[:2]
	printSlices("c", c)

	// you can see it, because cap of b is 5
	d := c[2:5]
	printSlices("d", d)
}

func printSlices(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
