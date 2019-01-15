package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map binds/maps keys to values
var m = map[string]Vertex{
	"Google":   {37.42, -122.08},
	"Facebook": {12.2, 23.43},
}
var s map[string]Vertex

func main() {
	s = make(map[string]Vertex)
	s["Bell Labs"] = Vertex{
		40.68, -74.39,
	}
	fmt.Println(s)
	fmt.Println(m)

	// insert, update, delete of an element in map d

	d := make(map[string]int)

	d["Answer"] = 42
	fmt.Println("The value:", d["Answer"])

	d["Answer"] = 48
	fmt.Println("The value:", d["Answer"])

	delete(d, "Answer")
	fmt.Println("The value:", d["Answer"])

	v, ok := d["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
