package client

import (
	"encoding/json"

	"github.com/adrianjjohnson/gomarvel"
)

// GetAllCharacters returns a list of character results
func (c *Client) GetAllCharacters() (gomarvel.Result, error) {
	uri := buildRoute(routePrefix, charactsers)
	params := addParameters(c.publicKey, c.privateKey)

	resp, err := c.client.Get(uri + params)
	if err != nil {
		return gomarvel.Result{}, err
	}
	defer resp.Body.Close()

	var result gomarvel.Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return gomarvel.Result{}, err
	}

	return result, nil
}
