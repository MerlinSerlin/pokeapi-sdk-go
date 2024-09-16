package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/merlinserlin/pokeapi-sdk-go/models" // Adjust this import path based on your project structure
)

func TestGeneration(t *testing.T) {
	// Mocked API response for Generation (can replace with actual response or mocked data)
	mockResponse := `{
		"id": 1,
		"name": "generation-i",
		"abilities": [
			{
				"name": "stench",
				"url": "https://pokeapi.co/api/v2/ability/1/"
			}
		],
		"main_region": {
			"name": "kanto",
			"url": "https://pokeapi.co/api/v2/region/1/"
		},
		"moves": [
			{
				"name": "pound",
				"url": "https://pokeapi.co/api/v2/move/1/"
			}
		],
		"pokemon_species": [
			{
				"name": "bulbasaur",
				"url": "https://pokeapi.co/api/v2/pokemon-species/1/"
			}
		],
		"types": [
			{
				"name": "normal",
				"url": "https://pokeapi.co/api/v2/type/1/"
			}
		],
		"version_groups": [
			{
				"name": "red-blue",
				"url": "https://pokeapi.co/api/v2/version-group/1/"
			}
		]
	}`

	var generation models.Generation
	err := json.Unmarshal([]byte(mockResponse), &generation)
	assert.NoError(t, err)

	// Assert top-level fields
	assert.Equal(t, 1, generation.ID)
	assert.Equal(t, "generation-i", generation.Name)

	// Assert Abilities
	assert.Len(t, generation.Abilities, 1)
	assert.Equal(t, "stench", generation.Abilities[0].Name)
	assert.Equal(t, "https://pokeapi.co/api/v2/ability/1/", generation.Abilities[0].Url)

	// Assert MainRegion
	assert.Equal(t, "kanto", generation.MainRegion.Name)
	assert.Equal(t, "https://pokeapi.co/api/v2/region/1/", generation.MainRegion.Url)

	// Assert Moves
	assert.Len(t, generation.Moves, 1)
	assert.Equal(t, "pound", generation.Moves[0].Name)
	assert.Equal(t, "https://pokeapi.co/api/v2/move/1/", generation.Moves[0].Url)

	// Assert PokemonSpecies
	assert.Len(t, generation.PokemonSpecies, 1)
	assert.Equal(t, "bulbasaur", generation.PokemonSpecies[0].Name)
	assert.Equal(t, "https://pokeapi.co/api/v2/pokemon-species/1/", generation.PokemonSpecies[0].Url)

	// Assert Types
	assert.Len(t, generation.Types, 1)
	assert.Equal(t, "normal", generation.Types[0].Name)
	assert.Equal(t, "https://pokeapi.co/api/v2/type/1/", generation.Types[0].Url)

	// Assert VersionGroups
	assert.Len(t, generation.VersionGroups, 1)
	assert.Equal(t, "red-blue", generation.VersionGroups[0].Name)
	assert.Equal(t, "https://pokeapi.co/api/v2/version-group/1/", generation.VersionGroups[0].Url)
}
