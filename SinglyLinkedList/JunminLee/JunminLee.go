package main

import "fmt"

type node struct {
	data     int
	nextNode *node
}

type linkdedList struct {
	head   *node
	length int
}

func (l *linkdedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.nextNode = second
	l.length++
}

// printListData - will print all nodes data
// it is a value reveiver (not aa pointer receiver)
// because we will print nodes data, we wont make a
// change to the reveiver
func (l linkdedList) printListData() {
	current := l.head
	for l.length != 0 {
		fmt.Printf("%d ", current.data)
		current = current.nextNode
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkdedList) deleteWithValue(value int) {
	if l.length == 0 {
		fmt.Println("empty LinkedList! Cannot delete!")
		return
	}
	if l.head.data == value {
		l.head = l.head.nextNode
		l.length--
		return
	}
	previous := l.head
	for previous.nextNode.data != value {
		if previous.nextNode.nextNode == nil {
			return
		}
		previous = previous.nextNode
	}
	previous.nextNode = previous.nextNode.nextNode
	l.length--

}

func main() {
	list := linkdedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	node4 := &node{data: 20}
	node5 := &node{data: 17}
	node6 := &node{data: 14}

	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.prepend(node5)
	list.prepend(node6)
	fmt.Println(list)

	fmt.Printf("\n")

	list.printListData()

	// list.deleteWithValue(1)
	// list.printListData()

	list.deleteWithValue(100)
	list.printListData()
	list.deleteWithValue(14)
	emptyList := linkdedList{}
	emptyList.deleteWithValue(1)
	list.deleteWithValue(14)
	list.printListData()
	list.deleteWithValue(17)
	list.printListData()
	list.deleteWithValue(1)
	list.printListData()
	list.deleteWithValue(3)
	list.printListData()
}
