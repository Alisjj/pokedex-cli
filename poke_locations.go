package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getClient(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("%v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *config) getLocations(url string) error {
	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &c.l); err != nil {
			return err
		}
		return nil
	}

	data, err := getClient(url)
	if err != nil {
		return err
	}
	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &c.l); err != nil {
		return err
	}

	return nil
}

func (c *config) exploreArea(area string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &c.c); err != nil {
			return err
		}
		return nil
	}

	data, err := getClient(url)

	if err != nil {
		return err
	}
	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &c.c); err != nil {
		return err
	}

	return nil
}

func (c *config) getPokemon(name string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	p := Pokemon{}

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &p); err != nil {
			return p, err
		}
		return p, nil
	}

	data, err := getClient(url)

	if err != nil {
		return p, err
	}
	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &p); err != nil {
		return p, err
	}

	return p, nil
}
