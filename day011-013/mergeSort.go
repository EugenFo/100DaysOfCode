package main

import "fmt"

func mergeSort(arr []int) []int {
	// checks if the array length is 1
	if len(arr) <= 1 {
		return arr
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
		rightArray = append(rightArray, arr[i])
	}

	// THIS WAS THE PROBLEM!!!!
	return merge(mergeSort(leftArray), mergeSort(rightArray))
}

// merge every halve of the array
func merge(leftArray, rightArray []int) []int {

	// creates empty slice
	result := []int{}

	leftLength := len(leftArray)
	rightLength := len(rightArray)

	// While Size of an Slice is greater than 0
	for leftLength > 0 && rightLength > 0 {

		// if first value of leftArray is less equal than first value of rightArray
		// then append leftArray value to result Slice and cut the appended value of
		// else append the rightArray value to result Slice
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

	// if length of leftArray is greater than 0
	// append every item of leftArray into result
	if leftLength > 0 {
		result = append(result, leftArray...)
	} else if rightLength > 0 {
		result = append(result, rightArray...)
	}
	return result
}

func main() {
	arr := []int{3, 6, 4, 8, 1, 2, 0, 10, 1}
	fmt.Println("Unsorted Slice:", arr)
	fmt.Println("Sorted Slice:", mergeSort(arr))
}
