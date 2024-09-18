package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Create a new cache instance
var c = cache.New(5*time.Minute, 10*time.Minute)

// Function to get data with caching
func GetCachedData(key string, fetchFunc func() (interface{}, error)) (interface{}, error) {
	if data, found := c.Get(key); found {
		return data, nil
	}

	data, err := fetchFunc()
	if err != nil {
		return nil, err
	}

	c.Set(key, data, cache.DefaultExpiration)
	return data, nil
}
