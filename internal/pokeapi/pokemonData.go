package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Cetch_Pokemmon(pokemon string) (PokemonData, error) {
	url := baseURL + "/pokemon/" + pokemon

	//if cache hit, return data directly
	if bytestream, hit := c.cache.Get(url); hit {
		pokemondata := PokemonData{}
		err := json.Unmarshal(bytestream, &pokemondata)
		if err != nil {
			return PokemonData{}, err
		}
		return pokemondata, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonData{}, err
	}

	pokemondata := PokemonData{}
	err = json.Unmarshal(dat, &pokemondata)
	if err != nil {
		return PokemonData{}, err
	}

	c.cache.Add(url, dat)
	
	return pokemondata, nil
}
