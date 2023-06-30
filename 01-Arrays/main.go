package main

import "fmt"

func main() {
	var myArray [4]string
	myArray[0] = "Apple"
	myArray[2] = "Cartot"
	fmt.Print("my Array is: ", myArray)
	fmt.Println(" , And my Array LENGTH is: ", len(myArray))

	var secArray = [5]int{44, 0, 8, 1}
	fmt.Println("my second Array is: ", secArray)

	tArry := [3]int{4, 5, 6}
	sum := 0
	for i := 0; i < len(tArry); i++ {
		// sum = sum + tArry[i]
		sum += tArry[i]
	}
	fmt.Println(sum)

	arr2D := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2D)

}
