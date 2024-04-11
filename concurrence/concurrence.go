package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutine:
	// Clients: en boucle
	//      - passent commande au serveur
	//		- boivent
	//		- attendent
	// Serveur:
	//		- reçoit des commandes de boisson
	//      - prépare la boisson et la renvoie au client
	// Patron:
	//		- attends 5 secondes
	//		- prévient tout le monde d'arrêter
	//
	// Hints:
	// - Lancer une goroutine: go func()
	// - Créer un channel: make(chan MyType)

	stopServeur := make(chan bool)
	stopClient := make(chan bool)
	commandes := make(chan Commande)

	go Serveur(commandes, stopServeur)

	for range 5 {
		go Client(commandes, stopClient)
	}

	time.Sleep(5 * time.Second)
	for range 5 {
		stopClient <- true
	}
	stopServeur <- true
}

type Commande struct {
	Boisson string
	Retour  chan string
}

func Serveur(commandes <-chan Commande, stop chan bool) {
	for {
		select {
		case commande := <-commandes:
			fmt.Println("Je prépare : " + commande.Boisson)
			commande.Retour <- commande.Boisson
		case <-stop:
			fmt.Println("J'arrête de bosser !")
			return
		}
	}
}

func Client(commandes chan<- Commande, stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("J'arrête de boire !")
			return
		default:
			retour := make(chan string)
			commande := Commande{
				Boisson: "Mojito",
				Retour:  retour,
			}
			commandes <- commande
			fmt.Println("J'ai commandé !")
			<-retour
			fmt.Println("Je bois !")
			time.Sleep(1 * time.Second)
		}
	}
}
