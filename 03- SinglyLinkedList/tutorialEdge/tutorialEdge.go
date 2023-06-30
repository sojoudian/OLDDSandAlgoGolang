package main

import "fmt"

// Node - a nore represent a link in our linked list
type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

// Insert -  add a new element to the start of our linked list
func (l *LinkedList) Insert(element int) {
	node := Node{
		Data: element,
		Next: l.Head,
	}
	l.Head = &node
	l.Length++
}

//DeleteFirst - removes the first element from our linked list
func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Length--
}

//List - iterates through all of the elements in our linked list and prints them
func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}

// Search - searches through every element of our linked list
// and returns the first element that matches otherwise it returns
// nil
func (l *LinkedList) Search(element int) *Node {
	current := l.Head
	for current != nil {
		if current.Data == element {
			return current
		}
		current = current.Next

	}
	return nil
}

// Delete - removes an element from Linked List if it matches the value
func (l *LinkedList) Delete(element int) {
	previous := l.Head
	current := l.Head
	for current != nil {
		if current.Data == element {
			previous.Next = current.Next
			l.Length--
		}
		previous = current
		current = current.Next
	}
}

func main() {
	fmt.Println("Singly Linked List")
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(4)
	ll.Insert(54)
	ll.Insert(3)
	ll.Insert(19)
	ll.Insert(17)
	// ll.List()
	//ll.DeleteFirst()
	// ll.List()

	ll.Delete(23)
	ll.List()
	fmt.Println(ll.Length)

	// if el := ll.Search(4); el != nil {
	// 	fmt.Printf("%+v\n", el)
	// }

}
