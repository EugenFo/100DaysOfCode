package main

import "fmt"

func mergeSort(arr []int) {
	// checks if the array length is 1
	if len(arr) <= 1 {
		return
	}
	// creates array for left and right side
	leftArray := []int{}
	rightArray := []int{}
	middle := len(arr)

	// appends left array values to the leftArray
	for i := 0; i < middle/2; i++ {
		leftArray = append(leftArray, arr[i])
	}

	// appends right array values to the rightArray
	for i := (middle / 2); i < middle; i++ {
		// fmt.Println(i)
		// fmt.Println(len(arr))
		rightArray = append(rightArray, arr[i])
	}
	fmt.Println("Before recursion", leftArray, rightArray)
	mergeSort(leftArray)
	mergeSort(rightArray)
	fmt.Println("after recursion", leftArray, rightArray)

	merge(leftArray, rightArray)
}

// merge every halve of the array
func merge(leftArray, rightArray []int) (result []int) {

	// creates empty slice
	result = []int{}

	// size := len(leftArray) + len(rightArray)
	leftLength := len(leftArray)
	rightLength := len(rightArray)

	for leftLength > 0 && rightLength > 0 {
		if leftArray[0] <= rightArray[0] {
			result = append(result, leftArray[0])
			leftArray = leftArray[1:]
			leftLength--
		} else {
			result = append(result, rightArray[0])
			rightArray = rightArray[1:]
			rightLength--
		}

	}
	if leftLength > 0 {
		result = append(result, leftArray...)
	} else if rightLength > 0 {
		result = append(result, rightArray...)
	}
	fmt.Println(result)
	return
}

func main() {
	arr := []int{3, 6, 4, 8, 1, 2, 0}
	mergeSort(arr)
}
