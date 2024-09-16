package api

import (
	"fmt"

	"github.com/merlinserlin/pokeapi-sdk-go/models"
	"github.com/merlinserlin/pokeapi-sdk-go/utils"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient *utils.HttpClient
}

func NewClient() *Client {
	return &Client{
		httpClient: utils.NewHttpClient(BaseURL),
	}
}

func (c *Client) GetPokemon(idOrName string) (*models.Pokemon, error) {
	url := fmt.Sprintf("/pokemon/%s", idOrName)
	pokemon := &models.Pokemon{}
	err := c.httpClient.Get(url, pokemon)
	if err != nil {
		return nil, fmt.Errorf("failed to get Pokemon: %w", err)
	}
	return pokemon, nil
}

func (c *Client) GetGeneration(idOrName string) (*models.Generation, error) {
	url := fmt.Sprintf("/generation/%s", idOrName)
	generation := &models.Generation{}
	err := c.httpClient.Get(url, generation)
	if err != nil {
		return nil, fmt.Errorf("failed to get Generation: %w", err)
	}
	return generation, nil
}
