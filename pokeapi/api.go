package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type locationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name *string `json:"name"`
		URL  *string `json:"url"`
	} `json:"results"`
}

var currentLocationArea locationArea

func GetPreviousLocationArea(client *Client) (locationArea, error) {
	err := currentLocationArea.getPreviousLocationArea(client)
	return currentLocationArea, err
}

func GetNextLocationArea(client *Client) (locationArea, error) {
	if currentLocationArea.Next == nil && currentLocationArea.Previous == nil {
		currentLocationArea.getNewLocationArea(client)
		return currentLocationArea, nil
	} else {
		err := currentLocationArea.getNextLocationArea(client)
		return currentLocationArea, err
	}
}

func (la *locationArea) getNewLocationArea(client *Client) {

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area", nil)
	res, err := client.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)

	if err = res.Body.Close(); err != nil {
		panic(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &la); err != nil {
		panic(err)
	}
	client.cache.Add("https://pokeapi.co/api/v2/location-area", body)
}

func (la *locationArea) getNextLocationArea(client *Client) error {
	if currentLocationArea.Next == nil {
		return errors.New("cannot get the previous location area")
	}
	next := *(currentLocationArea.Next)
	if v, found := client.cache.Get(next); found {
		if err := json.Unmarshal(v, &la); err != nil {
			return err
		}
	}

	req, err := http.NewRequest("GET", next, nil)
	if err != nil {
		return err
	}
	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)

	if err = res.Body.Close(); err != nil {
		return err
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &la); err != nil {
		return err
	}
	client.cache.Add(next, body)
	return nil
}

func (la *locationArea) getPreviousLocationArea(client *Client) error {
	if currentLocationArea.Previous == nil {
		return errors.New("cannot get the previous location area")
	}

	previous := *(currentLocationArea.Previous)
	if v, found := client.cache.Get(previous); found {
		if err := json.Unmarshal(v, &la); err != nil {
			return err
		}
	}

	req, err := http.NewRequest("GET", previous, nil)
	if err != nil {
		return err
	}
	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)

	if err = res.Body.Close(); err != nil {
		return err
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &la); err != nil {
		return err
	}
	client.cache.Add(previous, body)
	return nil
}
