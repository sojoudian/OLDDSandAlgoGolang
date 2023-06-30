package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Elements []int
}

// Push - the method which handles pushing new elements on top
// of our stack
func (s *Stack) Push(i int) {
	s.Elements = append(s.Elements, i)
}

//pop
func (s *Stack) Pop() (int, error) {
	if len(s.Elements) == 0 {
		return 0, errors.New("Stack is empty")
	} else {
		l := len(s.Elements) - 1
		var i int
		i = s.Elements[l]
		s.Elements = s.Elements[:l]
		return i, nil
	}
}

//Peek
// Peek - returns the top element from the stack without updating the stack
func (s *Stack) Peek() (int, error) {
	if len(s.Elements) == 0 {
		return 0, errors.New("Stack is empty")
	} else {
		peek := s.Elements[len(s.Elements)-1]
		return peek, nil
	}
}

// func (s Stack) Peek() (int, error) {
// 	if len(s.Elements) == 0 {
// 		return 0, errors.New("Stack is Empty")
// 	} else {
// 		v := s.Elements[len(s.Elements)-1]
// 		return v, nil
// 	}
// }

//IsEmpty
func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

//Length
func (s *Stack) Length() int {
	return len(s.Elements)
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(41)
	stack.Push(25)
	fmt.Println(stack)

	elem, _ := stack.Pop()
	fmt.Printf("element %d was deleted \n", elem)
	fmt.Println(stack)

	i, _ := stack.Peek()
	fmt.Println(i)
	fmt.Println(stack.Length())

}
