package main

import "fmt"

func main() {
	fmt.Println("Statick arrays")
	var myArray [8]int
	fmt.Println(myArray)

	mySecondArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(mySecondArray)

	firstElement_2ndArray := mySecondArray[0]
	lastElement_2ndArray := mySecondArray[len(mySecondArray)-1]
	fmt.Printf("The first element of the 2nd array is: %v, and the last element of 2nd array is: %v\n", firstElement_2ndArray, lastElement_2ndArray)

	// to print all elements of an array we can do the following loop
	for index, value := range mySecondArray {
		fmt.Printf("index is %v, and the value is %v\n", index, value)

	}

}
