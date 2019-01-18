package main

import (
	"bufio"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You are visiting the main page!</h1>")
}

func testPage1(w http.ResponseWriter, r *http.Request) {
	// get the path from the URL of the request
	path := html.EscapeString(r.URL.Path)
	fmt.Fprintf(w, "You are on: %q", path)
}

func imagePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Image Page")
}

func image(w http.ResponseWriter, r *http.Request) {
	// opens a jpg file
	f, _ := os.Open("write here a location of a image on the filesystem")
	// read jpg into memory
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	// set the content type header
	w.Header().Set("Content-Type", "image/jpeg")
	// write image to response
	w.Write(content)
}

func main() {
	// router
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/test1", testPage1)
	http.HandleFunc("/images", imagePage)
	http.HandleFunc("/image", image)

	// start server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
