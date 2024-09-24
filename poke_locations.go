package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *config) getLocations(url string) error {
	if val, ok := c.cache.Get(url); ok {

		if err := json.Unmarshal(val, &c.l); err != nil {
			return err
		}
		return nil
	}

	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	data, err := io.ReadAll(resp.Body)
	c.cache.Add(url, data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &c.l); err != nil {
		return err
	}

	return nil
}
