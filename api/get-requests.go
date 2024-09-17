package api

import (
	"fmt"
	"time"

	"github.com/merlinserlin/pokeapi-sdk-go/models"
	"github.com/merlinserlin/pokeapi-sdk-go/utils"
	"github.com/patrickmn/go-cache"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient *utils.HttpClient
	cache      *cache.Cache
}

func NewClient() *Client {
	return &Client{
		httpClient: utils.NewHttpClient(BaseURL),
		cache:      cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *Client) GetPokemon(idOrName string) (*models.Pokemon, error) {
	key := fmt.Sprintf("pokemon_%s", idOrName)

	data, err := utils.GetCachedData(key, func() (interface{}, error) {
		url := fmt.Sprintf("/pokemon/%s", idOrName)
		pokemon := &models.Pokemon{}
		err := c.httpClient.Get(url, pokemon)
		if err != nil {
			return nil, fmt.Errorf("failed to get Pokemon: %w", err)
		}
		return pokemon, nil
	})

	if err != nil {
		return nil, err
	}

	return data.(*models.Pokemon), nil
}

func (c *Client) GetGeneration(idOrName string) (*models.Generation, error) {
	key := fmt.Sprintf("generation_%s", idOrName)

	data, err := utils.GetCachedData(key, func() (interface{}, error) {
		url := fmt.Sprintf("/generation/%s", idOrName)
		generation := &models.Generation{}
		err := c.httpClient.Get(url, generation)
		if err != nil {
			return nil, fmt.Errorf("failed to get Generation: %w", err)
		}
		return generation, nil
	})

	if err != nil {
		return nil, err
	}

	return data.(*models.Generation), nil
}
