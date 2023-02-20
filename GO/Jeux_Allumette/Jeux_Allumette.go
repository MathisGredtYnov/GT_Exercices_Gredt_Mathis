package main

import (
	"fmt"
)

func JeuxDesAlumettes() {
	var joueur = 1
	var nbAlumette int
	var nbAllumetteRetire int
	fmt.Println("Bienvenue dans le jeux des allumettes !")
	fmt.Println("Avec combien d'allumette voulez-vous jouer ?")
	fmt.Scanln(&nbAlumette)
	for nbAlumette > 0 {
		fmt.Println("\033[H\033[2J")
		fmt.Println("Il reste", nbAlumette, "allumette(s)")
		fmt.Println("joueur", joueur, "Combien d'allumette voulez-vous retirer ?")
		fmt.Scanln(&nbAllumetteRetire)
		for nbAllumetteRetire <= 0 || nbAllumetteRetire > 3 {
			fmt.Println("Vous ne pouvez pas retirer", nbAllumetteRetire, "allumette(s)")
			fmt.Scanln(&nbAllumetteRetire)
		}
		nbAlumette -= nbAllumetteRetire
		if nbAlumette <= 0 {
			fmt.Println("Il reste 0 allumette")
			fmt.Println("joueur", joueur, "a perdu !")
		}
		if joueur == 1 {
			joueur = 2
		} else {
			joueur = 1
		}
	}
}

func main() {
	JeuxDesAlumettes()
}
