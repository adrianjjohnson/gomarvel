package gomarvel

import (
	"encoding/json"
	"fmt"

	"github.com/adrianjjohnson/gomarvel/model"
)

// GetComicByCharacterID returns a specific comic by characterID.
func (c *Client) GetComicByCharacterID(id int64) (model.ComicResults, error) {
	uri := fmt.Sprintf("%s%s%s/%d/%s?", baseURI, routePrefix, charactsers, id, comics)
	params := addParameters(c.publicKey, c.privateKey)

	resp, err := c.client.Get(uri + params)
	if err != nil {
		return model.ComicResults{}, err
	}
	defer resp.Body.Close()

	var result model.ComicResults
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return model.ComicResults{}, err
	}

	return result, nil
}
