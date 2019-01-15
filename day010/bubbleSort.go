package main

import (
	"fmt"
)

func bubbleSort(d []int) {
	unSorted := false
	for !unSorted {
		unSorted = true
		for i := 0; i < len(d)-1; i++ {
			if d[i] > d[i+1] {
				x := d[i]
				y := d[i+1]
				d[i] = y
				d[i+1] = x
				unSorted = false
			}
		}
	}
}

func main() {
	d := []int{54, 7, 39, 10, 54, 34, 2345, 343, 1, 300}
	fmt.Println("Unsorted Slice:", d)
	bubbleSort(d)
	fmt.Println("Sorted Slice:", d)
}
