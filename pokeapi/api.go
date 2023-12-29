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

func GetPreviousLocationArea() (locationArea, error) {
	err := currentLocationArea.getPreviousLocationArea()
	return currentLocationArea, err
}

func GetNextLocationArea() (locationArea, error) {
	if currentLocationArea.Next == nil && currentLocationArea.Previous == nil {
		currentLocationArea.getNewLocationArea()
		return currentLocationArea, nil
	} else {
		err := currentLocationArea.getNextLocationArea()
		return currentLocationArea, err
	}
}

func (la *locationArea) getNewLocationArea() {

	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
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
}

func (la *locationArea) getNextLocationArea() error {
	if currentLocationArea.Next == nil {
		return errors.New("cannot get the previous location area")
	}
	res, err := http.Get(*currentLocationArea.Next)
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
	return nil
}

func (la *locationArea) getPreviousLocationArea() error {
	if currentLocationArea.Previous == nil {
		return errors.New("cannot get the previous location area")
	}
	res, err := http.Get(*currentLocationArea.Previous)
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
	return nil
}
