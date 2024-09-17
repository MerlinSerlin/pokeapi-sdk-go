package tests

import (
	"testing"

	"github.com/merlinserlin/pokeapi-sdk-go/api"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemon(t *testing.T) {
	client := *api.NewClient()
	pokemon, err := client.GetPokemon("clefairy")

	if assert.NoError(t, err) {
		assert.NotNil(t, pokemon)
		assert.Equal(t, 35, pokemon.ID) // ID of Pikachu
		assert.Equal(t, "clefairy", pokemon.Name)
	}
}
