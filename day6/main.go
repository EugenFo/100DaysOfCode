package main

import "fmt"

type Rock struct {
	Mass   int
	Volume int
}

// Method to calculate the density of rocks + automatically implements Dense interface
func (r Rock) Density() int {
	fmt.Println(r)
	return r.Mass / r.Volume
}

// value arguments of Interface Type
// gets values passed by structs with an implemented Dense Interface
// Func doesn't care anymore what type of object it receives as long as the object implements the methods required by the interface
func IsItDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

type Geode struct {
}

// method to return density of geodes + automatically implements Dense interface, cause the interface
func (g Geode) Density() int {
	return 100
}

// creates interface called Dense for the Density Method.
//
type Dense interface {
	Density() int
}

func main() {
	a := Rock{Mass: 15, Volume: 10}
	// b := Rock{Mass: 10, Volume: 15}
	c := Geode{}
	fmt.Println(IsItDenser(a, c))
}
