package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Pages struct {
	Page1 string
	Page2 string
	Page3 string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl1 := template.Must(template.ParseFiles("index.html"))
	Pages := Pages{
		Page1: r.FormValue("page1"),
		Page2: r.FormValue("page2"),
		Page3: r.FormValue("page3"),
	}
	redirectToPage(w, r)
	tmpl1.Execute(w, Pages)

}

func redirectToPage(w http.ResponseWriter, r *http.Request) {
	// Permet de rediriger vers une autre page suivant le boutton cliqué //
	if r.FormValue("page") == "Page1" {
		http.Redirect(w, r, "/Page1", 301)
		return
	}
	// Permet de rediriger vers une autre page suivant le boutton cliqué //
	if r.FormValue("page") == "Page2" {
		http.Redirect(w, r, "/Page2", 301)
		return
	}
	// Permet de rediriger vers une autre page suivant le boutton cliqué //
	if r.FormValue("page") == "Page3" {
		http.Redirect(w, r, "/Page3", 301)
		return
	}
}

func main() {
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//gestion html
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Page1", Page1)
	http.HandleFunc("/Page2", Page2)
	http.HandleFunc("/Page3", Page3)
	http.ListenAndServe(":80", nil)
}

func Page1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Page1" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl1 := template.Must(template.ParseFiles("Page1.html"))
	Pages := Pages{}
	tmpl1.Execute(w, Pages)
}

func Page2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Page2" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl1 := template.Must(template.ParseFiles("Page2.html"))
	Pages := Pages{}
	tmpl1.Execute(w, Pages)
}

func Page3(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Page3" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl1 := template.Must(template.ParseFiles("Page3.html"))
	Pages := Pages{}
	tmpl1.Execute(w, Pages)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Rick Roll'd!")
	}
}
