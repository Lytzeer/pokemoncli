package main

import (
	"fmt"
	"math/rand"
)

// global variable

var pokemons = map[string]map[string]interface{}{
	"Bulbasaur": {
		"Type":  []string{"Grass", "Poison"},
		"HP":    22,
		"MaxHP": 22,
		"Dmg":   4,
	},
	"Charmander": {
		"Type":  []string{"Fire"},
		"HP":    24,
		"MaxHP": 24,
		"Dmg":   6,
	},
	"Squirtle": {
		"Type":  []string{"Water"},
		"HP":    23,
		"MaxHP": 23,
		"Dmg":   5,
	},
	"Pikachu": {
		"Type":  []string{"Electric"},
		"HP":    25,
		"MaxHP": 25,
		"Dmg":   5,
	},
	"Weedle": {
		"Type":  []string{"Bug", "Poison"},
		"HP":    16,
		"MaxHP": 16,
		"Dmg":   2,
	},
	"Ratata": {
		"Type":  []string{"Normal"},
		"HP":    20,
		"MaxHP": 20,
		"Dmg":   4,
	},
	"Sandshrew": {
		"Type":  []string{"Ground"},
		"HP":    19,
		"MaxHP": 19,
		"Dmg":   5,
	},
	"Nidoran": {
		"Type":  []string{"Poison"},
		"HP":    20,
		"MaxHP": 20,
		"Dmg":   4,
	},
	"Vulpix": {
		"Type":  []string{"Fire"},
		"HP":    23,
		"MaxHP": 23,
		"Dmg":   6,
	},
	"Psyduck": {
		"Type":  []string{"Water"},
		"HP":    17,
		"MaxHP": 23,
		"Dmg":   6,
	},
	"Mankey": {
		"Type":  []string{"Fighting"},
		"HP":    19,
		"MaxHP": 19,
		"Dmg":   7,
	},
	"Abra": {
		"Type":  []string{"Psychic"},
		"HP":    20,
		"MaxHP": 20,
		"Dmg":   6,
	},
	"Geodude": {
		"Type":  []string{"Rock", "Ground"},
		"HP":    23,
		"MaxHP": 23,
		"Dmg":   4,
	},
	"Magnemite": {
		"Type":  []string{"Electric", "Steel"},
		"HP":    24,
		"MaxHP": 24,
		"Dmg":   5,
	},
	"Mr. Mime": {
		"Type":  []string{"Psychic", "Fairy"},
		"HP":    24,
		"MaxHP": 24,
		"Dmg":   5,
	},
	"Eevee": {
		"Type":  []string{"Normal"},
		"HP":    20,
		"MaxHP": 20,
		"Dmg":   4,
	},
	"Vaporeon": {
		"Type":  []string{"Water"},
		"HP":    21,
		"MaxHP": 21,
		"Dmg":   4,
	},
	"Jolteon": {
		"Type":  []string{"Electric"},
		"HP":    23,
		"MaxHP": 23,
		"Dmg":   5,
	},
	"Flareon": {
		"Type":  []string{"Fire"},
		"HP":    22,
		"MaxHP": 22,
		"Dmg":   6,
	},
	"Sylveon": {
		"Type":  []string{"Fairy"},
		"HP":    21,
		"MaxHP": 21,
		"Dmg":   5,
	},
	"Espeon": {
		"Type":  []string{"Psychic"},
		"HP":    23,
		"MaxHP": 23,
		"Dmg":   4,
	},
	"Leafeon": {
		"Type":  []string{"Grass"},
		"HP":    22,
		"MaxHP": 22,
		"Dmg":   4,
	},
	"Glaceon": {
		"Type":  []string{"Ice"},
		"HP":    23,
		"MaxHP": 23,
		"Dmg":   5,
	},
	"Umbreon": {
		"Type":  []string{"Dark"},
		"HP":    25,
		"MaxHP": 25,
		"Dmg":   7,
	},
	"Mew": {
		"Type":  []string{"Psychic"},
		"HP":    27,
		"MaxHP": 27,
		"Dmg":   8,
	},
	"Jynx": {
		"Type":  []string{"Ice", "Psychic"},
		"HP":    20,
		"MaxHP": 20,
		"Dmg":   6,
	},
	"Glalie": {
		"Type":  []string{"Ice"},
		"HP":    19,
		"MaxHP": 19,
		"Dmg":   4,
	},
	"Ponyta": {
		"Type":  []string{"Fire"},
		"HP":    22,
		"MaxHP": 22,
		"Dmg":   5,
	},
	"Horsea": {
		"Type":  []string{"Water"},
		"HP":    21,
		"MaxHP": 21,
		"Dmg":   4,
	},
	"Vanillite": {
		"Type":  []string{"Ice"},
		"HP":    18,
		"MaxHP": 18,
		"Dmg":   4,
	},
	"Mareep": {
		"Type":  []string{"Electric"},
		"HP":    20,
		"MaxHP": 20,
		"Dmg":   5,
	},
	"Grimer": {
		"Type":  []string{"Poison"},
		"HP":    27,
		"MaxHP": 27,
		"Dmg":   6,
	},
	"Koffing": {
		"Type":  []string{"Poison"},
		"HP":    23,
		"MaxHP": 23,
		"Dmg":   4,
	},
	"Voltrob": {
		"Type":  []string{"Electric"},
		"HP":    21,
		"MaxHP": 21,
		"Dmg":   4,
	},
	"Magikarp": {
		"Type":  []string{"Water"},
		"HP":    14,
		"MaxHP": 14,
		"Dmg":   3,
	},
	"Cubchoo": {
		"Type":  []string{"Ice"},
		"HP":    18,
		"MaxHP": 18,
		"Dmg":   4,
	},
	"Torchic": {
		"Type":  []string{"Fire"},
		"HP":    19,
		"MaxHP": 19,
		"Dmg":   5,
	},
}
var potions = map[string]map[string]interface{}{
	"Potion": {
		"Description": "Heal 6 HP",
		"Heal":        6,
		"DropRate":    60,
	},
	"Super Potion": {
		"Description": "Heal 10 HP",
		"Heal":        10,
		"DropRate":    20,
	},
	"Hyper Potion": {
		"Description": "Heal 20 HP",
		"Heal":        20,
		"DropRate":    10,
	},
	"Full Restore": {
		"Description": "Restore all HP and remove any status",
		"Heal":        50,
		"DropRate":    5,
	},
	"Revive": {
		"Description": "Revive a Pokemon with half his life",
		"Heal":        0,
		"DropRate":    5,
	},
}

// player variable

var pokemon_inventory map[string]map[string]interface{} = make(map[string]map[string]interface{})
var potions_inventory []string = []string{}
var curent_pokemon string
var effect [][]string = [][]string{} // [0] = pokemon, [1] = effect, [2] = turns left, [3] = dmg

// ia variable

var pokemon_inventory_ia map[string]map[string]interface{} = make(map[string]map[string]interface{})
var potions_inventory_ia []string = []string{}

// player struct

type Player struct {
	name           string
	poke_inventory map[string]map[string]interface{}
	pot_inventory  []string
}

// main function

func main() {
	choosePokemon()
	choosePotions()
	player1 := Player{"Lytzeer", pokemon_inventory, potions_inventory}
	fmt.Println(player1.name)
	fmt.Println(getPokemonList())
	fmt.Println(getPotionList())
	name := getPokemonList()[0]
	fmt.Println("First Player pokemon name", name)
	fmt.Println("First Player pokemon type", getPokemonType(name))
	choosePokemonIA()
	choosePotionsIA()
}

// init player function

func choosePokemon() {
	key_pokemon := []string{}
	for k := range pokemons {
		key_pokemon = append(key_pokemon, k)
	}
	for i := 0; i < 5; i++ {
		pokemon_inventory[key_pokemon[i]] = pokemons[key_pokemon[i]]
		if i == 0 {
			curent_pokemon = key_pokemon[i]
		}
	}
}

func choosePotions() {
	key_potions := []string{}
	potions_droprate_list := []string{}
	for k := range potions {
		key_potions = append(key_potions, k)
	}
	for i := 0; i < 5; i++ {
		drop_rate := potions[key_potions[i]]["DropRate"].(int)
		for j := 0; j < drop_rate; j++ {
			potions_droprate_list = append(potions_droprate_list, key_potions[i])
		}
	}
	for i := 0; i < 7; i++ {
		random_potion := rand.Intn(len(potions_droprate_list)) + 1
		potions_inventory = append(potions_inventory, potions_droprate_list[random_potion])
	}
}

// init IA function

func choosePokemonIA() {
	key_pokemon := []string{}
	for k := range pokemons {
		key_pokemon = append(key_pokemon, k)
	}
	for i := 0; i < 5; i++ {
		pokemon_inventory_ia[key_pokemon[i]] = pokemons[key_pokemon[i]]
	}
	// fmt.Println("IA pokemon inventory :", pokemon_inventory_ia)
}

func choosePotionsIA() {
	key_potions := []string{}
	potions_droprate_list := []string{}
	for k := range potions {
		key_potions = append(key_potions, k)
	}
	for i := 0; i < 5; i++ {
		drop_rate := potions[key_potions[i]]["DropRate"].(int)
		for j := 0; j < drop_rate; j++ {
			potions_droprate_list = append(potions_droprate_list, key_potions[i])
		}
	}
	for i := 0; i < 7; i++ {
		random_potion := rand.Intn(len(potions_droprate_list)) + 1
		potions_inventory_ia = append(potions_inventory_ia, potions_droprate_list[random_potion])
	}
	// fmt.Println("IA backpack :", potions_inventory_ia)
}

// player get function

func getPokemonList() []string {
	pokemon_list := []string{}
	for k := range pokemon_inventory {
		pokemon_list = append(pokemon_list, k)
	}
	return pokemon_list
}

func getPotionList() []string {
	return potions_inventory
}

func getCurentPokemon() string {
	return curent_pokemon
}

func getPokemonType(pokemon_name string) []string {
	return pokemon_inventory[pokemon_name]["Type"].([]string)
}

// player combat function

func Attack(pokemon_inventory, pokemon_inventory_ia map[string]map[string]interface{}, name, ia_name string) {
	if pokemon_inventory_ia[ia_name]["HP"].(int)-pokemon_inventory[name]["Dmg"].(int) > 0 {
		pokemon_inventory_ia[ia_name]["HP"] = pokemon_inventory_ia[ia_name]["HP"].(int) - pokemon_inventory[name]["Dmg"].(int)
	} else {
		pokemon_inventory_ia[ia_name]["HP"] = 0
	}
}

func usePotion(pokemon_inventory, potions_inventory map[string]map[string]interface{}, pokemon_name, potion_name string) {
	if pokemon_inventory[pokemon_name]["HP"].(int)+potions[potion_name]["Heal"].(int) < pokemon_inventory[pokemon_name]["MaxHP"].(int) {
		pokemon_inventory[pokemon_name]["HP"] = pokemon_inventory[pokemon_name]["HP"].(int) + potions[potion_name]["Heal"].(int)
		fmt.Println("Your pokemon", pokemon_name, "has been healed by", potions[potion_name]["Heal"].(int), "HP")
	} else if pokemon_inventory[pokemon_name]["HP"].(int)+potions[potion_name]["Heal"].(int) > pokemon_inventory[pokemon_name]["MaxHP"].(int) {
		pokemon_inventory[pokemon_name]["HP"] = pokemon_inventory[pokemon_name]["MaxHP"].(int)
		fmt.Println("Your pokemon", pokemon_name, "has been healed by", (potions[potion_name]["Heal"].(int)+pokemon_inventory[pokemon_name]["HP"].(int))-pokemon_inventory[pokemon_name]["MaxHP"].(int), "HP")
	} else {
		fmt.Println("Your pokemon", pokemon_name, "is already full life")
	}
}

func changePokemon(name string) {
	curent_pokemon = name
	fmt.Println("Your current pokemon is now :", curent_pokemon)
}

// effects function

func Burn(pokemon_name string) {
	random := rand.Intn(100) + 1
	if random <= 10 {
		fmt.Println("Your pokemon", pokemon_name, "is burned")
		effect = append(effect, []string{pokemon_name, "Burn", "3", "3"})
	}
}

func Poison(pokemon_name string) {
	random := rand.Intn(100) + 1
	if random <= 15 {
		fmt.Println("Your pokemon", pokemon_name, "is poisoned")
		effect = append(effect, []string{pokemon_name, "Poison", "2", "5"})
	}
}

func Frozen(pokemon_name string) {
	random := rand.Intn(100) + 1
	if random <= 30 {
		fmt.Println("Your pokemon", pokemon_name, "is frozen")
		effect = append(effect, []string{pokemon_name, "Frozen", "1", "0"})
	}
}

func Paralysis(pokemon_name string) {
	random := rand.Intn(100) + 1
	if random <= 13 {
		fmt.Println("Your pokemon", pokemon_name, "is paralyzed")
		effect = append(effect, []string{pokemon_name, "Paralysis", "4", "0"})
	}
}

func Sleep(pokemon_name string) {
	random := rand.Intn(100) + 1
	if random <= 20 {
		fmt.Println("Your pokemon", pokemon_name, "is sleeping")
		effect = append(effect, []string{pokemon_name, "Sleep", "2", "0"})
	}
}
