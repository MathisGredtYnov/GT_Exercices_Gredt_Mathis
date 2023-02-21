package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func MenuFichier() {
	var boucle = false
	var choix int
	fmt.Println("Menu des actions sur le fichiers")
	fmt.Println()
	fmt.Println("1 - Récupérer tout le texte du fichier")
	fmt.Println("2 - Ajoutez du texte dans fichier")
	fmt.Println("3 - Supprimez le contenu du fichier")
	fmt.Println("4 - Remplacez le contenu du fichier par un autre text")
	fmt.Println("5 - Quitter")
	fmt.Scan(&choix)
	if boucle == false {
		if choix != 1 && choix != 2 && choix != 3 && choix != 4 && choix != 5 {
			fmt.Println("\033[H\033[2J")
			fmt.Println("Erreur de saisie")
			MenuFichier()
		}
		switch choix {
		case 1:
			fmt.Println("\033[H\033[2J")
			fmt.Println("Contenu du fichier : ")
			data, err := ioutil.ReadFile("text.txt")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(data))
			fmt.Println()
			fmt.Println()
			MenuFichier()
		case 2:
			var texte string
			fmt.Println("\033[H\033[2J")
			fmt.Print("Ajoutez du texte dans fichier : ")
			fmt.Scan(&texte)
			file, err := os.OpenFile("text.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
			_, err = file.WriteString(texte)
			if err != nil {
				panic(err)
			}
			file.Close()
			fmt.Println()
			fmt.Println()
			MenuFichier()
		case 3:
			fmt.Println("\033[H\033[2J")
			fmt.Print("tout le contenu du fichier à été supprimé")
			file, err := os.OpenFile("text.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
			_, err = file.WriteString("")
			if err != nil {
				panic(err)
			}
			file.Close()
			fmt.Println()
			fmt.Println()
			MenuFichier()
		case 4:
			var texte string
			fmt.Println("\033[H\033[2J")
			fmt.Print("Remplacez le contenu du fichier par un autre text : ")
			fmt.Scan(&texte)
			file, err := os.OpenFile("text.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
			_, err = file.WriteString(texte)
			if err != nil {
				panic(err)
			}
			file.Close()
			fmt.Println()
			fmt.Println()
			MenuFichier()
		case 5:
			boucle = true
		}
	}
}

func main() {
	MenuFichier()
}
