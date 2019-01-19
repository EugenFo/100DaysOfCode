package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// io library expects a byte slice

type Page struct {
	Title string
	Body  []byte
}

// global variable to cache templates
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// global variable
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//method for persistent storage

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// writes file to disk with 0600 permission
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}

// func that loads a textfile on the filesystem and represents it in the browser
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// validates the title
	/* title, err := getTitle(w, r)
	if err != nil {
		return
	} */
	p, err := loadPage(title)

	// if the viewed page don't exists, then it will redirected to the edit pag with all
	// the request and response; and a 302 statusCode (redirect)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// func to edit the textfile
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// validates the title
	/* title, err := getTitle(w, r)
	if err != nil {
		return
	} */

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// func to save the edit textfile
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// validates the title
	/* title, err := getTitle(w, r)
	if err != nil {
		return
	} */

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// func to render the template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// wrapper func that calls other func
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	/* 	p1 := &Page{Title: "Testpage", Body: []byte("Simple page")}
	   	p1.save()
	   	p2, _ := loadPage("Testpage")
	   	fmt.Println(string(p2.Body))
	*/
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
