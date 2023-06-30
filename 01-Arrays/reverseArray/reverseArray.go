package main

import "fmt"

func reverseArraybyCopying(a []int32) []int32 {

	return a
}

func main() {
	myArr := [...]int{1, 4, 3, 2}
	b := myArr
	size := len(myArr)
	fmt.Println("Printing the original array:", myArr)
	for i := 0; i < size; i++ {
		myArr[i] = b[size-i-1]
	}
	fmt.Println("Printing the reversed array:", myArr)
}
