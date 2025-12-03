package pokeapi

import (
	"encoding/json"
	"io"

	//"io"
	"net/http"
)

func (c *Client) LocationExplore(location string) (DeepMapData, error) {
	url := baseURL + "/location-area/" + location

	//if cache hit, return data directly
	if bytestream, hit := c.cache.Get(url); hit {
		mapdata := DeepMapData{}
		err := json.Unmarshal(bytestream, &mapdata)
		if err != nil {
			return DeepMapData{}, err
		}
		return mapdata, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DeepMapData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return DeepMapData{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return DeepMapData{}, err
	}

	mapData := DeepMapData{}
	err = json.Unmarshal(dat, &mapData)
	if err != nil {
		return DeepMapData{}, err
	}

	c.cache.Add(url, dat)
	return mapData, nil

}
