package main

import (
	"fmt"
	"log"

	"github.com/merlinserlin/pokeapi-sdk-go/api"
)

func main() {
	// Initialize the client from the api package
	client := api.NewClient()

	// Get a Pokémon
	pokemon, err := client.GetPokemon("pikachu")
	if err != nil {
		log.Fatalf("Error getting Pokemon: %v", err)
	}
	fmt.Printf("Pokemon: %s, ID: %d\n", pokemon.Name, pokemon.ID)

	// Get a Generation
	gen, err := client.GetGeneration("generation-i")
	if err != nil {
		log.Fatalf("Error getting Generation: %v", err)
	}
	fmt.Printf("Generation: %s, Main Region: %s\n", gen.Name, gen.MainRegion.Name)
}
