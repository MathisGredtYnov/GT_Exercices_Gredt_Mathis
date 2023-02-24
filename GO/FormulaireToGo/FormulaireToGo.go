package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type FormulaireStruct struct {
	Nom           string
	Prenom        string
	Age           string
	DateNaissance string
}

var f FormulaireStruct

func (f *FormulaireStruct) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	details := FormulaireStruct{
		Nom:    r.FormValue("nom"),
		Prenom: r.FormValue("prenom"),
		Age:    r.FormValue("age"),
	}
	tmp.Execute(w, details)
	EcrireDonneFormulaire(w, r)
}

func EcrireDonneFormulaire(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Nom : ", r.FormValue("nom"))
	fmt.Println("Prenom : ", r.FormValue("prenom"))
	fmt.Println("Age : ", r.FormValue("age"))
}

func main() {
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//gestion html
	http.HandleFunc("/", f.indexHandler)
	http.ListenAndServe(":8080", nil)
}
