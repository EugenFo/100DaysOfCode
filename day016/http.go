package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// download from a website
	resp, _ := http.Get("https://google.de/asdwef")

	if strings.Contains(resp.Status, "200") {
		fmt.Println("Download was successful")
	} else {
		fmt.Println("Download was not successful")
	}

	// close the body after everything is done
	defer resp.Body.Close()

	// read everything from body
	body, _ := ioutil.ReadAll(resp.Body)

	// converts bytes into strings
	bodyText := string(body)

	fmt.Println("Done")
	fmt.Println("response length:", len(bodyText))

	runeSlice := []rune(bodyText)
	fmt.Println("First Chars:", string(runeSlice))

}
