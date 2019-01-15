package main

import "fmt"

// struct == collation of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{3, 223})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	p := &v // p is going to be a pointer to v
	fmt.Println("value of p:", *p)
	fmt.Println("memory address of p:", &p)
	fmt.Printf("memory address of v: %p\n", &v)
	fmt.Printf("memory address of p: %p\n", &p)
	fmt.Printf("memory address of v: %p\n", p)

	p.X = 100 // (*p).X is also possible
	fmt.Println("value of v", v)
	fmt.Printf("Print value of p: %v\n", *p)
}
