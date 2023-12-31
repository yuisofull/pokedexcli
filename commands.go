package main

import (
	"errors"
	"fmt"
	"github.com/yuisofull/pokedexcli/pokeapi"
	"math/rand"
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
		"catch": {
			name:        "catch",
			description: "catch a pokemon. Usage: catch <pokemon_name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a pokemon. Usage: inspect <pokemon_name>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all pokemons you have caught",
			callback:    commandPokedex,
		},
	}
}

func commandPokedex(cfg *config, param ...string) error {
	fmt.Println("Your Pokedex: ")
	for pkm := range cfg.PokeDex {
		fmt.Println(" - " + pkm)
	}
	return nil
}

func commandInspect(cfg *config, param ...string) error {
	if len(param) == 0 {
		return errors.New("missing 1 argument. Usage: inspect <pokemon_name>")
	}
	if pkm, exist := cfg.PokeDex[param[0]]; !exist {
		fmt.Println("you have not caught that pokemon")
		return nil
	} else {
		fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats: 
`, pkm.Name, pkm.Height, pkm.Weight)
		for _, stat := range pkm.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, ty := range pkm.Types {
			fmt.Printf("  -%s\n", ty.Type.Name)
		}
	}
	return nil
}

func commandCatch(cfg *config, param ...string) error {
	if len(param) == 0 {
		return errors.New("missing 1 argument. Usage: catch <pokemon_name>")
	}
	if _, exist := cfg.PokeDex[param[0]]; exist {
		fmt.Println("you have already caught it!")
		return nil
	}
	pkm, err := pokeapi.GetPokemon(&cfg.pokeapiClient, param[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", param[0])
	r := rand.Intn(pkm.BaseExperience)

	if r > 40 {
		fmt.Printf("%s escaped!\n", param[0])
		return nil
	}
	fmt.Printf("%s was caught!\n"+
		"You may now inspect it with the inspect command.\n", param[0])
	cfg.PokeDex[pkm.Name] = *pkm
	return nil
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
