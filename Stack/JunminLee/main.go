package main

import "fmt"

type Stack struct {
	items []int
}

//Push - will add a value at the top of the stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop - will remove a value at the top of the stack
// And returns the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

//Peek
func (s Stack) Peek(i int) {
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(1)
	myStack.Push(34)
	myStack.Push(21)
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)
}
