package main

import (
	"fmt"
)

func stringToSlice(s string) []string {
	var sl []string
	for _, j := range s {
		sl = append(sl, string(j))
	}
	return sl
}

func isBalanced(s string) bool {
	sl := stringToSlice(s)
	var x int
	var y int

	for i := range sl {
		if string(sl[i]) == "x" {
			x++
		} else if string(sl[i]) == "y" {
			y++
		}
	}

	if x == y {
		return true
	}

	return false

}

func main() {
	fmt.Println(isBalanced("xxxyyy"))
	fmt.Println(isBalanced("yyyxxx"))
	fmt.Println(isBalanced("xxxyyyy"))
	fmt.Println(isBalanced("yyxyxxyxxyyyyxxxyxyx"))
	fmt.Println(isBalanced("xyxxxxyyyxyxxyxxyy"))
	fmt.Println(isBalanced(""))
	fmt.Println(isBalanced("x"))
}

//https://www.reddit.com/r/dailyprogrammer/comments/afxxca/20190114_challenge_372_easy_perfectly_balanced/
