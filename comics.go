package gomarvel

import (
	"encoding/json"
	"fmt"

	"github.com/adrianjjohnson/gomarvel/model"
)

// GetAllComics returns a list of character results
func (c *Client) GetAllComics() (model.ComicResult, error) {
	uri := fmt.Sprintf("%s%s%s?", baseURI, routePrefix, charactsers)
	params := addParameters(c.publicKey, c.privateKey)

	resp, err := c.client.Get(uri + params)
	if err != nil {
		return model.ComicResult{}, err
	}
	defer resp.Body.Close()

	var result model.ComicResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return model.ComicResult{}, err
	}

	return result, nil
}

// GetComicByCharacterID returns a specific comic by characterID.
func (c *Client) GetComicByCharacterID(id int64) (model.ComicResult, error) {
	uri := fmt.Sprintf("%s%s%s/%d/%s?", baseURI, routePrefix, charactsers, id, comics)
	params := addParameters(c.publicKey, c.privateKey)

	resp, err := c.client.Get(uri + params)
	if err != nil {
		return model.ComicResult{}, err
	}
	defer resp.Body.Close()

	var result model.ComicResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return model.ComicResult{}, err
	}

	return result, nil
}
