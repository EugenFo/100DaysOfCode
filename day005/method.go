package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method == function with a special receiver
// receiver is declared between func keyword and the func name
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer receiver, changes the value of Vertex
// a value receiver would not change the value of Vertex, cause it's would be only a copy of Vertex
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// same func like Scale, but it's not a method, just a regular function, who has also a pointer receiver
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	/* v := Vertex{3, 4}
	fmt.Println("before Scale:", v)
	// change the value of v
	v.Scale(10)
	fmt.Println("after scale:", v)
	fmt.Println(v.Abs()) */

	a := Vertex{3, 4}
	b := &a
	a.Scale(2)
	fmt.Println("memory address of a:", &b)
	// &a, to transmit a pointer and not a value
	ScaleFunc(&a, 10)

	// create a pointer to Vertex{3, 4}
	p := &Vertex{3, 4}
	p.Scale(3)
	fmt.Println("memory address of p", &p)
	// transmit a pointer to the func
	ScaleFunc(p, 8)

	fmt.Println(a, p)

	// method with value receivers takes either a value or a pointer
	// methods with pointer receivers takes either a value or a pointer

	// functions with a value argument, only takes a value; pointer would create a compile error
	// functions with a pointer argument, only takes a pointer; value would create a compile error

	// 2 reasons to use pointer receiver
	// 1: modify the value that its receiver points to
	// 2: avoid copy of the value on each method call. Good, if the struct is very large
}
