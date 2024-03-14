package main

import (
	"fmt"
	"slices"
)

func main() {

	var legumes = make([]string, 0)
	legumes = append(legumes, "salade", "tomate", "oignon")
	autre := []string{"champignon", "poids", "carottes"}
	legumes = append(legumes, autre...)
	slices.Delete(legumes, 3, 4)
	// slices.DeleteFunc(legumes, del)

	fmt.Printf("%#v", legumes)

}

func del(elem string) bool {
	return elem != "poids"
}
