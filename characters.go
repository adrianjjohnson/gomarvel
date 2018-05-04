package gomarvel

import (
	"encoding/json"

	"github.com/adrianjjohnson/gomarvel/model"
)

// GetAllCharacters returns a list of character results
func (c *Client) GetAllCharacters() (model.Result, error) {
	uri := buildRoute(routePrefix, charactsers)
	params := addParameters(c.publicKey, c.privateKey)

	resp, err := c.client.Get(uri + params)
	if err != nil {
		return model.Result{}, err
	}
	defer resp.Body.Close()

	var result model.Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return model.Result{}, err
	}

	return result, nil
}
