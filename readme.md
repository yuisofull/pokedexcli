# Pokedex CLI

Welcome to the Pokedex Command Line Interface (CLI) project! This CLI application allows you to interact with the PokeAPI to explore Pokemon data, catch new Pokemon, and manage your Pokedex.

## Features

- **Interactive Command Line Interface**: Use commands to explore and interact with the Pokemon world.
- **Pokedex Management**: Keep track of the Pokemon you've caught in your Pokedex.
- **Location Exploration**: Explore different areas to discover Pokemon encounters.
- **Pokemon Inspection**: Get detailed information about a specific Pokemon.

## Getting Started

1. **Installation**: Clone the repository to your local machine.

   ```bash
   git clone https://github.com/your-username/pokedex-cli.git
   ```

2. **Dependencies**: Ensure you have Go installed. If not, you can download it [here](https://golang.org/dl/).

3. **Run the Application**:

   ```bash
   cd pokedex-cli
   go run main.go
   ```

4. **Explore and Enjoy!**: Use the available commands to navigate the Pokemon world.

## Commands

- `help`: Display a help message with available commands.
- `map`: Display the next names of 20 location areas.
- `mapb`: Display the previous names of 20 location areas.
- `explore <area_name>`: Display Pokemon names in a specific location area.
- `catch <pokemon_name>`: Catch a Pokemon.
- `inspect <pokemon_name>`: Inspect details of a caught Pokemon.
- `pokedex`: List all Pokemon you have caught.

## Dependencies

- [PokeAPI](https://pokeapi.co/): Utilized for retrieving Pokemon and location information.
- [Pokecache](https://github.com/yuisofull/pokedexcli/pokecache): Custom cache module for efficient data storage.
