package main

type Animal interface {
	Crier()
}

type Elephant struct{}

func (h *Elephant) Crier() {}

type Canard struct{}

func (i Canard) Crier()  {}
func (i Canard) Dormir() {}

type Chimère struct {
	// Méthode Crier sur
	// les deux composants
	Elephant
	Canard
}

func main() {
	// var c Animal
	an := GetAnimaux()
	for _, animal := range an {
		animal.Crier()
	}
	// Lever l’ambiguïté
	//c.Elephant.Crier()
	//c.Canard.Crier()
}

func GetAnimaux() []Animal {
	return []Animal{
		&Canard{},
		&Elephant{},
	}

}
