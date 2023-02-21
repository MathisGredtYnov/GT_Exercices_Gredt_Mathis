package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type tableau struct {
	Objet []int
}

func TableauDimension(w http.ResponseWriter, r *http.Request) []int {
	var max int = 25
	var min int = 3
	var tableau []int
	rand.Seed(time.Now().UnixNano())
	NombreDeDimension := rand.Intn(max-min) + min
	for i := 0; i < NombreDeDimension; i++ {
		tableau = append(tableau, rand.Int())
	}
	return tableau
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	details := tableau{
		Objet: TableauDimension(w, r),
	}
	tmp.Execute(w, details)
}

func main() {
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//gestion html
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
