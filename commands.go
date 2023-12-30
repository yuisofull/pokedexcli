package main

import (
	"fmt"
	"github.com/yuisofull/pokedexcli/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
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
	}
}

func commandMapBack(cfg *config) error {
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

func commandMap(cfg *config) error {
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

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
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
