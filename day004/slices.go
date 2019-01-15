package main

import "fmt"

func main() {
	/* // array, cause it has a fixed size
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// slice, cause it is dynamically sized
	b := []int{0, 1, 2, 3}
	fmt.Println(b)

	// b[1:3] includes 1, excludes 3
	var s []int = b[1:3]
	fmt.Println(s) */
	// _________________________________________
	/* // example for length and capacity of a slice
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s) */
	// _____________________________________________
	/* // built-in append function

	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1, 2)
	printSlice(s)

	s = append(s, 4, 5, 6)
	printSlice(s) */
	// ______________________________-

	// range iterates over a slice or map
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// i is the index of the map; v contains the value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// prints only the index and omits the value
	for i := range pow {
		fmt.Println(i)
	}
	// prints only the value and omits the index; _, omits index
	for _, value := range pow {
		fmt.Println(value)
	}
	// ________________________________

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
