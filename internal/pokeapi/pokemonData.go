package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonData(pokemon string) (pokemonData, error) {
	url := baseURL + "/pokemon/" + pokemon

	//if cache hit, return data directly
	if bytestream, hit := c.cache.Get(url); hit {
		pokemondata := pokemonData{}
		err := json.Unmarshal(bytestream, &pokemondata)
		if err != nil {
			return pokemonData{}, err
		}
		return pokemondata, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemonData{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemonData{}, err
	}

	pokemondata := pokemonData{}
	err = json.Unmarshal(dat, &pokemondata)
	if err != nil {
		return pokemonData{}, err
	}

	c.cache.Add(url, dat)
	return pokemondata, nil
}
