package main

import (
	"strings"
	"unicode"
)

func scoreCalc(s string) {
	r := []rune(s)
	score := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0}
	for i := 0; i < len(s); i++ {
		if unicode.IsLower(r[i]) == true {
			score[strings.ToLower(string(s[i]))]++

		} else {
			score[strings.ToLower(string(s[i]))]--
		}
	}

	// perform the sorting
	sorting(score)
}

/* func sorting(m map[string]int) {
	fmt.Println(m)
	sa := map[int]string{}
	sl := []int{}
	for k, v := range m {
		sa[v] = k
		sl = append(sl, v)
	}
	fmt.Println(sa, sl)
	sort.Ints(sl)
	fmt.Println("after sort:", sl)

	for _, v := range sl {
		fmt.Println(sa[v], v)
	}
} */

func main() {
	scoreCalc("dbbaCEDbdAacCEAadcB")
}

// https://www.reddit.com/r/dailyprogrammer/comments/8jcffg/20180514_challenge_361_easy_tally_program/
