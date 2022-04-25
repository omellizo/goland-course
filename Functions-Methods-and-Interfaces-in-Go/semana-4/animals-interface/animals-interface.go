package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (cow Cow) Eat() {
	fmt.Println(cow, "grass")
}

func (cow Cow) Move() {
	fmt.Println(cow, "walk")
}

func (cow Cow) Speak() {
	fmt.Println(cow, "moo")
}

type Bird struct {
	name string
}

func (bird Bird) Eat() {
	fmt.Println(bird, "worms")
}

func (bird Bird) Move() {
	fmt.Println(bird, "fly")
}

func (bird Bird) Speak() {
	fmt.Println(bird, "peep")
}

type Snake struct {
	name string
}

func (snake Snake) Eat() {
	fmt.Println(snake, "mice")
}

func (snake Snake) Move() {
	fmt.Println(snake, "slither")
}

func (snake Snake) Speak() {
	fmt.Println(snake, "hsss")
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

func newAnimal(nameAnimale string, typeAnimal string) Animal {
	switch typeAnimal {
	case "cow":
		return Cow{nameAnimale}
	case "bird":
		return Bird{nameAnimale}
	case "snake":
		return Snake{nameAnimale}
	default:
		fmt.Println("no case")
		return nil
	}
}

func informationRequested(animal Animal, infoRequest string) {
	switch infoRequest {
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	animals := make(map[string]Animal)

	for {
		// newanimal anyNameAnimal  “cow”, “bird”, or “snake”
		// query nameAnimal “eat”, “move”, or “speak”
		request := getString("Request", scanner)
		inputs := strings.Split(request, " ")

		switch inputs[0] {
		case "newanimal":
			animal := newAnimal(inputs[1], inputs[2])
			animals[inputs[1]] = animal
			fmt.Println("Created it!")
		case "query":
			animal := animals[inputs[1]]
			informationRequested(animal, inputs[2])
		default:
			fmt.Println("no case")
		}
	}
}
