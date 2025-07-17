package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationList(pageURL *string) (ShallowMapData, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShallowMapData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ShallowMapData{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ShallowMapData{}, err
	}

	mapData := ShallowMapData{}
	err = json.Unmarshal(dat, &mapData)
	if err != nil {
		return ShallowMapData{}, err
	}

	return mapData, nil
}
