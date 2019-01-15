package main

import "fmt"

type dog struct {
	name  string
	breed string
}

// Dog's Methods
func (d dog) walk() {
	fmt.Println(d.name, "walks across the room")
}

func (d dog) sit() {
	fmt.Println(d.name, "sits down")
}

func (d dog) fetch() {
	fmt.Println(d.name, "fetches a toy")
}

type cat struct {
	name  string
	breed string
}

// Cat's Methods
func (c cat) walk() {
	fmt.Println(c.name, "walks across the room")
}

func (c cat) sit() {
	fmt.Println(c.name, "sits down")
}

func (c cat) purr() {
	fmt.Println(c.name, "purrs")
}

// Functions
// Don't need it anymore, cause are better for this
/* func demoDog(dog Dog) {
	dog.Walk()
	dog.Sit()
}

func demoCat(cat Cat) {
	cat.Walk()
	cat.Sit()
} */

func demo(animal fourLegged) {
	animal.walk()
	animal.sit()
}

// Interface for the same methods for Cat and Dog

type fourLegged interface {
	walk()
	sit()
}

func main() {
	dog := dog{"Rex", "Schäferhund"}
	cat := cat{"Mietze", "Katze"}
	/* 	demoDog(dog)
	demoCat(cat) */
	demo(dog)
	demo(cat)
}

/*
	demoDog and demoCat Functions are the same, would be nice if we could create a function that takes both arguments of Dog and Cat
	=> less maintenance, cause only 1 Function and not 2
*/

// interface I embeds k(), j() and i(), because it embeds interface J, which embeds interface K
/*
type I interface {
	J
	i()
}

type J interface {
	K
	j()
}

type K interface {
	k()
}
*/
