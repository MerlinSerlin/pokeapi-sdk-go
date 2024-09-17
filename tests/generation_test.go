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
		assert.NotNil(t, generation)
		assert.Equal(t, 1, generation.ID) // Generation ID for Generation I
		assert.Equal(t, "generation-i", generation.Name)
		assert.NotNil(t, generation.MainRegion)
		assert.NotNil(t, generation.Abilities)
		assert.NotNil(t, generation.Moves)
		assert.NotNil(t, generation.PokemonSpecies)
		assert.NotNil(t, generation.Types)
		assert.NotNil(t, generation.VersionGroups)
	}
}
