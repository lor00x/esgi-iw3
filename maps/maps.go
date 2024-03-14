package main

import (
	"fmt"
	"maps"
)

func main() {
	var schools = map[string]string{
		"PARIS":      "ESGI",
		"NANTES":     "ESGI",
		"NICE":       "ESSI",
		"STRASBOURG": "???",
	}

	bis := map[string]string{
		"NICE":    "eirughizuegh",
		"ALBI":    "HOP",
		"BORDEAU": "ESGI",
	}

	maps.Copy(bis, schools)

	clear(bis)
	for town, school := range bis {
		fmt.Println(town, " -> ", school)
	}
}
