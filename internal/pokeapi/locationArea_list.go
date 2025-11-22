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

	var resp *http.Response
	var err error

	//if cache hit, return data directly
	if bytestream, hit := c.cache.Get(url); hit {
		mapdata := ShallowMapData{}
		err := json.Unmarshal(bytestream, &mapdata)
		if err != nil {
			return ShallowMapData{}, err
		}
		return mapdata, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShallowMapData{}, err
	}

	resp, err = c.httpClient.Do(req)
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

	c.cache.Add(url, dat)
	return mapData, nil
}
