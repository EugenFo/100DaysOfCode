package main

import (
	"fmt"
)

func main() {
	// create empty map with [string] as keys and string as values
	smth := map[string]string{}
	// fmt.Println(smth)
	smth = map[string]string{
		"bird":  "blue",
		"snake": "green",
		"cat":   "black",
	}
	// fmt.Println(smth)

	// Get the value of snake
	sn := smth["snake"]
	fmt.Println(sn)

	// add a key-value pair to a map
	smth["dog"] = "white"
	dg := smth["dog"]
	fmt.Println(dg)

	// delete of key-value pair
	fmt.Println("length before deleting:", len(smth))
	delete(smth, "snake")
	fmt.Println("length after deleting:", len(smth))

	// loop through the map

	for key, value := range smth {
		fmt.Println(key, "=", value)
	}

	// look up if a key exsists in the map; ok is often used for
	// the boolean value
	// ok == true can be written as just ok
	if exists, ok := smth["cat"]; ok == true {
		fmt.Println("cat", exists)
	}

	if exists, ok := smth["snake"]; ok {
		fmt.Println("snake", exists)
	} else {
		fmt.Println("snake does not exists in the map")
	}

	// copy only the key of a map into a slice
	slice := []string{}
	for key, _ := range smth {
		slice = append(slice, key)
	}
	fmt.Println(slice)

	// with make you can specify a capacity of a map, slice, array
	// this will optimize the map, slice, array and you can add items much faster in it
	// because you don't need to resize and reallocate the items in the map.
	/*
		es := make(map[int]string, 200)

		for i := 0; i <=200; i++{
			es[i] = "as[i]"
		}
		fmt.Println(es)

		// range keyword randomize the order of the keys
		for key, value := range es {
			fmt.Println(key, "=", value)
		}
	*/

	/*
		// capacity Performance
		t0 := time.Now()

		for i := 0; i < 2000; i++ {
			values := make(map[int]int, 1000)
			for x := 0; x < 1000; x++ {
				values[x] = x
			}
			if values[0] != 0 {
				fmt.Println(0)
			}
		}

		t1 := time.Now()

		for i := 0; i < 2000; i++ {
			values := map[int]int{}
			for x := 0; x < 1000; x++ {
				values[x] = x
			}
			if values[0] != 0 {
				fmt.Println(0)
			}
		}

		t2 := time.Now()

		fmt.Println(t1.Sub(t0))
		fmt.Println(t2.Sub(t1))
	*/

	// you can pass a map as an argument to a func
	printSmth(smth)
	
	// func to handle if a item already exists in the map
	// if yes then exit the func
	handled := map[string]bool{}
	test(handled, "bird")
	test(handled, "cat")
	test(handled, "bird")
	fmt.Println(handled)
}

func printSmth(c map[string]string) {
	fmt.Println("func:", c["bird"])
}


func test(handled map[string]bool, name string){
	// check if the value of the key name is true, then return the func
	if _, ok := handled[name]; ok{
		fmt.Println("ALREADY HANDLED:", name)
		return
	}
	// else add the name key with the value true into the map
	fmt.Println("ADDED:", name)
	handled[name] = true
}