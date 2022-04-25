package main

import "fmt"

// No checks for wrong inputs are being made as I didnt see those requirements in the instructions.
//
func main() {

	animalsCreated := make(map[string]Animal)
	for true {
		var command, animalName, animaltypeOrInformation string
		fmt.Println(">")
		fmt.Scanln(&command, &animalName, &animaltypeOrInformation)

		switch command {
		case "newanimal":
			//c := animalFactory(animaltypeOrInformation, animalName)
			animalsCreated[animalName] = animalFactory(animaltypeOrInformation, animalName)
			fmt.Println("Created it!")
		case "query":
			switch animaltypeOrInformation {
			case "eat":
				animalsCreated[animalName].Eat()
			case "move":
				animalsCreated[animalName].Move()
			case "speak":
				animalsCreated[animalName].Speak()
			}

		}
	}
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	Name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	Name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	Name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func animalFactory(animalType string, animalName string) Animal {
	var retAnimal Animal

	switch animalType {

	case "cow":
		retAnimal := Cow{Name: animalName}
		return retAnimal
	case "bird":
		retAnimal := Bird{Name: animalName}
		return retAnimal
	case "snake":
		retAnimal := Snake{Name: animalName}
		return retAnimal
	default:
		return retAnimal
	}
}
