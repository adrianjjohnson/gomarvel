package gomarvel

import (
	"encoding/json"
	"fmt"

	"github.com/adrianjjohnson/gomarvel/model"
)

// GetAllCharacters returns a list of character results
func (c *Client) GetAllCharacters() (model.CharacterResult, error) {
	uri := fmt.Sprintf("%s%s%s?", baseURI, routePrefix, charactsers)
	params := addParameters(c.publicKey, c.privateKey)

	resp, err := c.client.Get(uri + params)
	if err != nil {
		return model.CharacterResult{}, err
	}
	defer resp.Body.Close()

	var result model.CharacterResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return model.CharacterResult{}, err
	}

	return result, nil
}

// GetCharacterByID gets a single character
func (c *Client) GetCharacterByID(id int64) (model.CharacterResult, error) {
	uri := fmt.Sprintf("%s%s%s/%d?", baseURI, routePrefix, charactsers, id)
	params := addParameters(c.publicKey, c.privateKey)

	resp, err := c.client.Get(uri + params)
	if err != nil {
		return model.CharacterResult{}, err
	}
	defer resp.Body.Close()

	var result model.CharacterResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return model.CharacterResult{}, err
	}

	return result, nil
}
