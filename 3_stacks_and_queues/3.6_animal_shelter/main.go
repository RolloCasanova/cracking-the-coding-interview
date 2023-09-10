package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 1 {
		panic("usage: go run main.go [<dog|cat> <name>]...")
	}

	// creating the shelter from the arguments
	shelter := AnimalShelter(os.Args[1:])

	// to reuse these variables, we need to declare them outside the loop
	var (
		operation string
		animal    animal
	)

	for {
		// some operations don't require a value, so we only need to read the operation

		fmt.Println("enter operation to perform on shelter (enqueue, dequeue, dequeue_dog, dequeue_cat), or exit to quit:")
		_, err := fmt.Scan(&operation)
		if err != nil {
			break
		}

		// if the operation is enqueue, we need to read the animal kind and name
		if operation == "enqueue" {

			fmt.Println("enter animal kind: ")
			_, err := fmt.Scan(&animal.kind)
			if err != nil {
				panic("animal kind is required and must be either \"dog\" or \"cat\"")
			}

			if animal.kind != "dog" && animal.kind != "cat" {
				panic("animal kind must be either \"dog\" or \"cat\"")
			}

			fmt.Println("enter animal name: ")
			_, err = fmt.Scan(&animal.name)
			if err != nil {
				panic("animal name is required")
			}

			animal.timestamp = time.Now()
		}

		switch operation {
		case "enqueue":
			shelter.Enqueue(animal)
		case "dequeue":
			shelter.DequeueAny()
		case "dequeue_dog":
			shelter.DequeueDog()
		case "dequeue_cat":
			shelter.DequeueCat()
		case "print":
			shelter.Print()
		case "exit":
			return
		default:
			fmt.Println("invalid operation \" " + operation + "\"")
		}

		fmt.Println()
	}
}

func AnimalShelter(data []string) *shelter {
	// arguments are in the form <dog|cat> <name> (in pairs)
	if len(data)%2 != 0 {
		fmt.Println("usage: go run main.go [<dog|cat> <name>]...")

		return nil
	}

	shelter := &shelter{}
	for i := 0; i < len(data); i += 2 {
		if data[i] != "dog" && data[i] != "cat" {
			fmt.Printf("invalid animal kind: %s\n", data[i])

			return nil
		}

		shelter.Enqueue(animal{data[i], data[i+1], time.Now()})
	}

	return shelter
}

// animal is a struct that represents an animal in the shelter
type animal struct {
	kind      string
	name      string
	timestamp time.Time
}

// AnimalShelterQueue is an interface for an animal shelter queue
type AnimalShelterQueue interface {
	Enqueue(a animal)
	DequeueAny() *animal
	DequeueDog() *animal
	DequeueCat() *animal
}

// linkedlist is a linked list of animals
type linkedlist struct {
	value animal
	next  *linkedlist
}

// shelter is a struct that represents an animal shelter
type shelter struct {
	dogs *linkedlist
	cats *linkedlist
}

// Enqueue adds an animal to the shelter
func (s *shelter) Enqueue(animal animal) {
	switch animal.kind {
	case "dog":
		if s.dogs == nil {
			s.dogs = &linkedlist{animal, nil}
			return
		}

		dogsStart := s.dogs

		for s.dogs.next != nil {
			s.dogs = s.dogs.next
		}

		s.dogs.next = &linkedlist{animal, nil}
		s.dogs = dogsStart

	case "cat":
		if s.cats == nil {
			s.cats = &linkedlist{animal, nil}
			return
		}

		catsStart := s.cats

		for s.cats.next != nil {
			s.cats = s.cats.next
		}
		s.cats.next = &linkedlist{animal, nil}
		s.cats = catsStart

	default:
		fmt.Printf("invalid animal kind: %s", animal.kind)
	}
}

// DequeueAny removes the oldest animal from the shelter
func (s *shelter) DequeueAny() (animal animal) {
	dog := s.dogs
	cat := s.cats

	if dog == nil && cat == nil {
		fmt.Println("shelter is empty")

		return
	}

	if dog == nil {
		animal = cat.value
		s.cats = cat.next

		return
	}

	if cat == nil {
		animal = dog.value
		s.dogs = dog.next

		return
	}

	if dog.value.timestamp.Before(cat.value.timestamp) {
		animal = dog.value
		s.dogs = dog.next
	} else {
		animal = cat.value
		s.cats = cat.next
	}

	return
}

// DequeueDog removes the oldest dog from the shelter
func (s *shelter) DequeueDog() (animal animal) {
	if s.dogs == nil {
		fmt.Println("no dogs in shelter")

		return
	}

	animal = s.dogs.value
	s.dogs = s.dogs.next

	return animal
}

// DequeueCat removes the oldest cat from the shelter
func (s *shelter) DequeueCat() (animal animal) {
	if s.cats == nil {
		fmt.Println("no cats in shelter")
		return
	}

	animal = s.cats.value
	s.cats = s.cats.next

	return animal
}

// Print prints the shelter by arrival order
func (s *shelter) Print() {
	dogStart := s.dogs
	catStart := s.cats

	dog := s.dogs
	cat := s.cats

	for dog != nil && cat != nil {
		if dog.value.timestamp.Before(cat.value.timestamp) {
			fmt.Printf("%s %s - (%s)\n", dog.value.kind, dog.value.name, dog.value.timestamp)
			dog = dog.next
		} else {
			fmt.Printf("%s %s - (%s)\n", cat.value.kind, cat.value.name, cat.value.timestamp)
			cat = cat.next
		}
	}

	for dog != nil {
		fmt.Printf("%s %s - (%s)\n", dog.value.kind, dog.value.name, dog.value.timestamp)
		dog = dog.next
	}

	for cat != nil {
		fmt.Printf("%s %s - (%s)\n", cat.value.kind, cat.value.name, cat.value.timestamp)
		cat = cat.next
	}

	s.dogs = dogStart
	s.cats = catStart
}
