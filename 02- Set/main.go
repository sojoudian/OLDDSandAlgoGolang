package main

import (
	"errors"
	"fmt"
)

// Set struct - is our representation of a set data structure
type Set struct {
	Elements map[string]struct{}
}

func main() {
	mySet := NewSet()
	mySet.Add("Earth")
	mySet.Add("Venus")
	mySet.Add("Mars")
	mySet.Add("Earth")
	mySet.List()

	fmt.Printf("\n+++++++++++++++++++++++++++++\n\n")

	mySet.Delete("Venus")
	mySet.List()

	fmt.Printf("\n+++++++++++++++++++++++++++++\n\n")
	fmt.Println(mySet.Contains("Mars"))
	fmt.Println(mySet.Contains("Venus"))
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	fmt.Println("This is the new set", &set)
	return &set
}

// Add - adds element to our set data structure
func (s *Set) Add(element string) {
	s.Elements[element] = struct{}{}
}

// Delete - deletes an element from our set if it exists
func (s *Set) Delete(element string) error {
	if _, exists := s.Elements[element]; !exists {
		return errors.New("element is not present in the mentioned set data structure")
	}
	delete(s.Elements, element)
	return nil
}

// Contains - checks if an element exists in our set data structure
func (s *Set) Contains(element string) bool {
	_, exists := s.Elements[element]
	return exists
}

// List - lits all elements in our set data structure
func (s *Set) List() {
	for key, _ := range s.Elements {
		fmt.Println(key)
	}
}
