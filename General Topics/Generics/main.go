package main

import (
	"fmt"
)

type numbers interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

// func newGenericFunc[age int64 | float64](myAge age) {
// 	fmt.Println(myAge)
// }

// func newGenericFunc[age any](myAge age) {
// 	// fmt.Println(myAge)
// 	val := int(myAge) + 2
// 	fmt.Println(val)

// }

// func newGenericFunc[age int64 | float64](myAge age) {
// 	val := float64(myAge) + 2
// 	fmt.Println(val)
// }

func newGenericFunc[age numbers](myAge age) {
	val := float64(myAge) + 2
	fmt.Println(val)
}

func BubleSort[N numbers](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func main() {
	var testAge1 int64 = 33
	var testAge2 float64 = 33.2
	newGenericFunc(testAge1)
	newGenericFunc(testAge2)

	// var myAgeInString string = "Thirty three"
	//newGenericFunc(myAgeInString)

	list := []int64{4, 15, 17, 1, 6, 13, 3}
	listF := []float32{4.75, 15.76, 17.78, 1.79, 6.80, 13.81, 3.82}
	sorted1 := BubleSort(list)
	sorted2 := BubleSort(listF)
	fmt.Println("", sorted1)
	fmt.Println("", sorted2)

}
