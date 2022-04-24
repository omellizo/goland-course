package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func getString(label string, scanner *bufio.Scanner) string {
	var res string
	for {
		fmt.Printf("Enter %s: ", label)
		if scanner.Scan() {
			res = scanner.Text()
			break
		} else {
			fmt.Println("Invalid input.")
		}
	}
	return res
}

func loadData() map[string]Animal {
	animals := make(map[string]Animal)

	cow := Animal{"grass", "walk", "moo"}
	animals["cow"] = cow

	bird := Animal{"worms", "fly", "peep"}
	animals["bird"] = bird

	snake := Animal{"mice", "slither", "hsss"}
	animals["snake"] = snake

	return animals
}

func main() {
	animals := loadData()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		request := getString("Request", scanner)
		inputs := strings.Split(request, " ")
		animal := animals[inputs[0]]

		switch inputs[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("no case")
		}
	}

}
