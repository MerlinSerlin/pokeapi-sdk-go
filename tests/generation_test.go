package tests

import (
	"testing"

	"github.com/merlinserlin/pokeapi-sdk-go/api"
	"github.com/stretchr/testify/assert"
)

func TestGetGeneration(t *testing.T) {
	client := *api.NewClient()
	generation, err := client.GetGeneration("1")

	if assert.NoError(t, err) {
		// General assertions
		assert.NotNil(t, generation)
		assert.Equal(t, 1, generation.ID) // Generation ID for Generation I
		assert.Equal(t, "generation-i", generation.Name)

		// MainRegion assertions
		assert.NotNil(t, generation.MainRegion)
		assert.Equal(t, "kanto", generation.MainRegion.Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/region/1/", generation.MainRegion.Url)

		// Abilities assertions
		assert.NotNil(t, generation.Abilities)
		assert.Equal(t, 0, len(generation.Abilities)) // No abilities in Generation I

		// Moves assertions
		assert.NotNil(t, generation.Moves)
		assert.Greater(t, len(generation.Moves), 0) // Ensure there are moves
		assert.Equal(t, "pound", generation.Moves[0].Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/move/1/", generation.Moves[0].Url)

		// PokemonSpecies assertions
		assert.NotNil(t, generation.PokemonSpecies)
		assert.Greater(t, len(generation.PokemonSpecies), 0)
		assert.Equal(t, "bulbasaur", generation.PokemonSpecies[0].Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/pokemon-species/1/", generation.PokemonSpecies[0].Url)

		// Types assertions
		assert.NotNil(t, generation.Types)
		assert.Greater(t, len(generation.Types), 0)
		assert.Equal(t, "normal", generation.Types[0].Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/type/1/", generation.Types[0].Url)

		// VersionGroups assertions
		assert.NotNil(t, generation.VersionGroups)
		assert.Equal(t, 2, len(generation.VersionGroups)) // Generation I has 2 version groups
		assert.Equal(t, "red-blue", generation.VersionGroups[0].Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/version-group/1/", generation.VersionGroups[0].Url)
		assert.Equal(t, "yellow", generation.VersionGroups[1].Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/version-group/2/", generation.VersionGroups[1].Url)
	}
}
