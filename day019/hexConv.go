package main

import (
	"fmt"
)

func hexcolor(red, green, blue int) string {
	r := fmt.Sprintf("%02x", red)
	g := fmt.Sprintf("%02x", green)
	b := fmt.Sprintf("%02x", blue)
	return r + g + b
}

func main() {
	fmt.Print("#", hexcolor(255, 99, 71), "\n")
	fmt.Print("#", hexcolor(184, 134, 11), "\n")
	fmt.Print("#", hexcolor(189, 183, 107), "\n")
	fmt.Print("#", hexcolor(0, 0, 205), "\n")
}

//https://www.reddit.com/r/dailyprogrammer/comments/a0lhxx/20181126_challenge_369_easy_hex_colors/
