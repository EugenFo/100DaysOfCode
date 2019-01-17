package main

import (
	"fmt"
	"sort"
)

// creates a type with the methods
type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	// sorting in Go operates on slices
	sn := []string{"cat", "bird", "snake", "dog", "zebra", "elephant", "monkey", "bird", "fly"}
	fmt.Println("unsorted slice:", sn)
	// Strings() sorts the strings in the slice in a ascending order
	sort.Strings(sn)
	fmt.Println("sorted slice:", sn)

	// for the Sort() func we need to define an Interface with the methods Len, Less and Swap
	// it's not sorting in an alphabetic order
	newEl := []string{"ruby", "java", "python", "perl", "go", "c++", "haskell", "php", "javascript", "erlang", "c#", "f#", "bash", "hack", "assembler"}
	fmt.Println("before sorting:", newEl)
	sort.Sort(ByLen(newEl))
	fmt.Println("after sorting:", newEl)

	// assign the keys into a slice and sort the slice, then combine the value of the slice with the keys of
	// the map and represent keySlice = valueMap; so the map is getting sorted
	as := map[string]int{
		"xyz": 1,
		"abc": 1,
		"abd": 2,
		"fge": 1,
		"oef": 1,
	}

	sl := []string{}
	for key, _ := range as {
		sl = append(sl, key)
	}

	sort.Strings(sl)
	// loop over sorted key-value pair
	for i := range sl {
		key := sl[i]
		value := as[key]
		fmt.Printf("%v = %v\n", key, value)
	}

	// StringsAreSorted returns a bool if the strings in a slice are already sorted or not
	// empty slice == sorted
	animalsSorted := []string{"bird", "cat", "zebra"}
	empty := []string{}
	animalsUnsorted := []string{"cat", "bird"}

	if sort.StringsAreSorted(animalsSorted) {
		fmt.Println("animalsSorted are sorted:", animalsSorted)
	}

	if sort.StringsAreSorted(empty) {
		fmt.Println("empty is also sorted:", empty)
	}

	if sort.StringsAreSorted(animalsUnsorted) {
		fmt.Println("will never reach this point")
	} else {
		fmt.Println("animalsUnorted is not sorted:", animalsUnsorted)
	}

	// sort can of course also sort ints, but we must pass Ints() an int slice
	numbers := []int{5, 3, 4, 10, 200, 666, 342, 42}

	sort.Ints(numbers)
	fmt.Println(numbers)

	// IntsAreSorted works the same like StringsAreSorted
	numbersSorted := []int{1, 2, 3, 4, 5}

	if sort.IntsAreSorted(numbersSorted) {
		fmt.Println(numbersSorted)
	}
}
