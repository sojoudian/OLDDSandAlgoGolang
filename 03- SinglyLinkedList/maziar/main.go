package main

import "fmt"

// Node - a nore represent a link in our linked list
type Node struct {
	Data     int
	NextNode *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

// Prepend - always add a new element to the head of our linked list
func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.NextNode = second
	l.Length++
}

// Insert -  add a new element to the start of our linked list
func (l *LinkedList) Insert(value int) {
	node := Node{
		NextNode: l.Head,
		Data:     value,
	}
	l.Head = &node
	l.Length++
	l.PrintListData()
}

// DeleteFirst - removes the first element from our linked list
func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.NextNode
	l.Length--
	l.PrintListData()
}

// Size - returns the size of our linked list
func (l LinkedList) Size() {
	size := l.Length
	fmt.Printf("LinkedList size is %v \n", size)
}

// printListData - will print all nodes data
// it is a value reveiver (not aa pointer receiver)
// because we will print nodes data, we wont make a
// change to the reveiver
func (l LinkedList) PrintListData() {
	current := l.Head
	// l.Length != 0  is equal to
	//l.Head != nil
	for l.Length != 0 {
		fmt.Printf("%d ", current.Data)
		current = current.NextNode
		l.Length--
	}
	fmt.Printf("\n")
}

// List - iterates through all of the elements in our linked list and prints them
func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v =>> ", *current)
		current = current.NextNode
	}

}

// Search - searches through every element of our linked list
// and returns the first element that matches otherwise it returns
// nil
func (l *LinkedList) Search(value int) *Node {
	current := l.Head
	for current != nil {
		if current.Data == value {
			fmt.Printf("%d is find at %v \n", value, current)
			return current
		}
		current = current.NextNode
	}
	return nil
}

// Delete - removes an element from Linked List if it matches the value
func (l *LinkedList) Delete(value int) {
	fmt.Printf("Value is %d \n", value)
	if l.Length == 0 {
		fmt.Println("empty LinkedList!")
		return
	}
	if l.Head.Data == value {
		fmt.Printf("value %d removed \n", value)
		l.Head = l.Head.NextNode
		l.Length--
		return
	}
	previous := l.Head
	for previous.NextNode.Data != value {
		if previous.NextNode.NextNode == nil {
			fmt.Printf("couldn't find %d in the list, and %d is not deleted\n", value, value)
			return
		}
		previous = previous.NextNode
	}
	previous.NextNode = previous.NextNode.NextNode
	l.Length--
}

func (l *LinkedList) AddtoTail(value int) {
	current := l.Head
	for l.Length != 0 {
		if current.NextNode == nil {
			node := Node{
				NextNode: nil,
				Data:     value,
			}
			current.NextNode = &node
			l.Length++
			l.PrintListData()
			break
		}
		current = current.NextNode
	}
}

func main() {
	myList := LinkedList{}
	node1 := &Node{Data: 1}
	node2 := &Node{Data: 2}
	node3 := &Node{Data: 3}
	node4 := &Node{Data: 20}
	node5 := &Node{Data: 17}
	node6 := &Node{Data: 14}
	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)
	myList.Prepend(node4)
	myList.Prepend(node5)
	myList.Prepend(node6)
	myList.Delete(44444)

	fmt.Println("++++++++++++++++")
	myList.PrintListData()
	myList.List()
	fmt.Println("\n++++++++++++++++")
	// emptyList := LinkedList{}
	// emptyList.Delete(23)
	fmt.Printf("\n")

	myList.Delete(1)
	myList.PrintListData()
	myList.Delete(100)
	myList.PrintListData()
	myList.Delete(14)
	myList.PrintListData()
	myList.Delete(14)
	myList.PrintListData()
	myList.Delete(17)
	myList.PrintListData()
	myList.Delete(1)
	myList.PrintListData()
	myList.Delete(3)
	myList.PrintListData()

	myList.Insert(696)
	myList.Insert(1368)
	myList.List()
	myList.Search(20)
	myList.Size()

	myList.DeleteFirst()

	myList.AddtoTail(18)
	myList.List()
	myList.PrintListData()

	node7 := &Node{Data: 314}
	myList.Prepend(node7)
	myList.PrintListData()

}

// Delete			done
// Prepend			done
// PrintListData	done
