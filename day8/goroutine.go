package main

import "fmt"

/* func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
} */

func f(f string) {
	for i := 0; i < 3; i++ {
		fmt.Println(f, ":", i)
	}
}

func main() {
	/* go say("World")
	say("Hello") */

	// blocking function
	f("direct")

	// func that will be proceed concurrently
	go f("goroutine")

	// 2nd func that will be proceed concurrently
	go func(msg string) {
		fmt.Println(msg)
	}("2nd goroutine")

	//need userinput to stop the program
	fmt.Scanln()
	fmt.Println("done")

}
