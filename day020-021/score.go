package main

import (
	"fmt"
	"sort"
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

func sorting(m map[string]int) {
	type kv struct {
		key   string
		value int
	}

	var sa []kv

	for k, v := range m {
		sa = append(sa, kv{k, v})
	}

	sort.Slice(sa, func(i, j int) bool {
		return sa[i].value > sa[j].value
	})

	for _, kv := range sa {
		fmt.Printf("%s:%d, ", kv.key, kv.value)
	}
	fmt.Println("")
}

func main() {
	scoreCalc("abcde")
	scoreCalc("dbbaCEDbdAacCEAadcB")
	scoreCalc("EbAAdbBEaBaaBBdAccbeebaec")
}

// https://www.reddit.com/r/dailyprogrammer/comments/8jcffg/20180514_challenge_361_easy_tally_program/
