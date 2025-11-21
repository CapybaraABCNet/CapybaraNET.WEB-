package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type TheUser struct {
	Username string
	UserAge  int
}

type Capybara struct {
	Name  string
	Money int
}

var (
	capybara = Capybara{
		Name:  "Капибара",
		Money: 0,
	}
	mu sync.Mutex
)

var (
	ClickCount int
	Cen        int = 45
	Click      int = 1
	Text       string
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func m(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/m.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func prikol(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/prikol.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func abc(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/abc.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func progstat(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/progstat.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func science(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/science.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func shutk(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/shutk.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func i(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/i.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func fortune(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/fortune.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func WebSiteStart() {
	http.HandleFunc("/", index)
	http.HandleFunc("/m", m)
	http.HandleFunc("/progstat", progstat)
	http.HandleFunc("/science", science)
	http.HandleFunc("/shutk", shutk)

	http.HandleFunc("/prikol", prikol)
	http.HandleFunc("/i", i)
	http.HandleFunc("/fortune", fortune)
	http.HandleFunc("/abc", abc)

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	WebSiteStart()
	fmt.Println("Запуск сервера...")
	if err := http.ListenAndServe(":8637", nil); err != nil {
		panic(err)
	}
}
