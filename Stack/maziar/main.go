package main

import (
	"errors"
	"fmt"
)

//Stack - representing stack data structure
type Stack struct {
	items []int
}

// Length returns the size of the stack
func (s Stack) Length() int {
	return len(s.items)
}

func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(newVal int) {
	s.items = append(s.items, newVal)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		lastItemIndex := len(s.items) - 1
		lastItemVal := s.items[lastItemIndex]
		s.items = s.items[:lastItemIndex]
		return lastItemVal, nil
	}

}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		lastItemIndex := len(s.items) - 1
		lastItemVal := s.items[lastItemIndex]
		return lastItemVal, nil
	}
}

func main() {
	stack := Stack{}
	fmt.Println("length of stack is", stack.Length())
	fmt.Println("Is stack empty or not?", stack.IsEmpty())
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println("length of stack is", stack.Length())
	fmt.Println("Is stack empty or not?", stack.IsEmpty())
	fmt.Println(stack)
	i, _ := stack.Peek()
	fmt.Println("Peek of the stack is ", i)

	j, _ := stack.Pop()
	fmt.Println(j, "was poped from the stack")
	fmt.Println("length of stack is", stack.Length())
	fmt.Println("Is stack empty or not?", stack.IsEmpty())
	fmt.Println(stack)

}
