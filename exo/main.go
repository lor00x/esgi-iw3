package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	list := GenerateRandomStrings()
	list = RemoveStringsWithCharE(list)
	DisplayList(list)
	// TODO:
	// Génère une liste de 5 chaînes entre 3 et 10 char aléatoires

	// Ecrire une fonction pour filtrer la liste
	// et enlever les éléments qui contiennent la lettre E

	// Afficher la liste

}

func GenerateRandomStrings() []string {
	return []string{"rijnf", "ijef", "oijf", "fuh", "oeiujhf"}
}

func RemoveStringsWithCharE(list []string) []string {
	del := func(elem string) bool {
		return strings.Contains(strings.ToLower(elem), "e")
	}
	return slices.DeleteFunc(list, del)
}

func DisplayList(list []string) {
	fmt.Println(list)
}
