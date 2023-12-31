package main

import (
	"errors"
	"fmt"
	"github.com/yuisofull/pokedexcli/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, param ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the next names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous names of 20 location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "displays Pokemon's names in a location area. Usage: explore <area_name>",
			callback:    commandExplore,
		},
	}
}
func commandExplore(cfg *config, param ...string) error {
	if len(param) == 0 {
		return errors.New("missing 1 argument. Usage: explore <area_name>")
	}
	laInfo, err := pokeapi.GetLocationAreaInfo(&cfg.pokeapiClient, param[0])
	if err != nil {
		return err
	}
	fmt.Println(`Exploring ` + param[0] + `...
Found Pokemon: `)
	for _, pokemon := range laInfo.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}
func commandMapBack(cfg *config, param ...string) error {
	la, err := pokeapi.GetPreviousLocationArea(&cfg.pokeapiClient)
	if err != nil {
		//fmt.Println("Cannot get the previous location area")
		return err
	}
	for _, location := range la.Results {
		fmt.Println(*location.Name)
	}
	return nil
}

func commandMap(cfg *config, param ...string) error {
	la, err := pokeapi.GetNextLocationArea(&cfg.pokeapiClient)
	if err != nil {
		//fmt.Println("Cannot get the next location area")
		return err
	}
	for _, location := range la.Results {
		fmt.Println(*location.Name)
	}
	return nil
}

func commandExit(cfg *config, param ...string) error {
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, param ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
