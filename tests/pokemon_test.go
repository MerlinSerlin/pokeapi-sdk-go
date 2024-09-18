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
		assert.Equal(t, 35, pokemon.ID)
		assert.Equal(t, "clefairy", pokemon.Name)
		assert.Equal(t, 113, pokemon.BaseExperience)
		assert.Equal(t, 6, pokemon.Height)
		assert.True(t, pokemon.IsDefault)
		assert.Equal(t, 75, pokemon.Weight)

		// Test Abilities
		assert.Len(t, pokemon.Abilities, 3)
		assert.Equal(t, "cute-charm", pokemon.Abilities[0].Ability.Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/ability/56/", pokemon.Abilities[0].Ability.URL)
		assert.False(t, pokemon.Abilities[0].IsHidden)
		assert.Equal(t, 1, pokemon.Abilities[0].Slot)

		assert.Equal(t, "magic-guard", pokemon.Abilities[1].Ability.Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/ability/98/", pokemon.Abilities[1].Ability.URL)
		assert.False(t, pokemon.Abilities[1].IsHidden)
		assert.Equal(t, 2, pokemon.Abilities[1].Slot)

		assert.Equal(t, "friend-guard", pokemon.Abilities[2].Ability.Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/ability/132/", pokemon.Abilities[2].Ability.URL)
		assert.True(t, pokemon.Abilities[2].IsHidden)
		assert.Equal(t, 3, pokemon.Abilities[2].Slot)

		// Test Forms
		assert.Len(t, pokemon.Forms, 1)
		assert.Equal(t, "clefairy", pokemon.Forms[0].Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/pokemon-form/35/", pokemon.Forms[0].URL)

		// Test Game Indices
		assert.Len(t, pokemon.GameIndices, 20)
		assert.Equal(t, 4, pokemon.GameIndices[0].GameIndex)
		assert.Equal(t, "red", pokemon.GameIndices[0].Version.Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/version/1/", pokemon.GameIndices[0].Version.URL)

		// Test Held Items
		assert.Len(t, pokemon.HeldItems, 3)
		assert.Equal(t, "moon-stone", pokemon.HeldItems[0].Item.Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/item/81/", pokemon.HeldItems[0].Item.URL)
		assert.Len(t, pokemon.HeldItems[0].VersionDetails, 22)
		assert.Equal(t, 5, pokemon.HeldItems[0].VersionDetails[0].Rarity)
		assert.Equal(t, "ruby", pokemon.HeldItems[0].VersionDetails[0].Version.Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/version/7/", pokemon.HeldItems[0].VersionDetails[0].Version.URL)

		assert.Equal(t, "leppa-berry", pokemon.HeldItems[1].Item.Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/item/131/", pokemon.HeldItems[1].Item.URL)
		assert.Len(t, pokemon.HeldItems[1].VersionDetails, 12)
		assert.Equal(t, 50, pokemon.HeldItems[1].VersionDetails[0].Rarity)

		assert.Equal(t, "comet-shard", pokemon.HeldItems[2].Item.Name)
		assert.Equal(t, "https://pokeapi.co/api/v2/item/624/", pokemon.HeldItems[2].Item.URL)
		assert.Len(t, pokemon.HeldItems[2].VersionDetails, 4)
		assert.Equal(t, 1, pokemon.HeldItems[2].VersionDetails[0].Rarity)

		// Test Location Area Encounters
		assert.Equal(t, "https://pokeapi.co/api/v2/pokemon/35/encounters", pokemon.LocationAreaEncounters)

		// The rest of the tests go on...
		// (Species, Sprites, Stats, Types, PastTypes)
	}
}
