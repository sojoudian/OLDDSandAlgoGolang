package main

import (
	"fmt"
	"sort"
)

func main() {
	mySlice := []string{"hi"}
	fmt.Println(mySlice)

	var myThirdSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(myThirdSlice)

	var myFourthSlice = []int{}
	fmt.Print(myFourthSlice, " \t")
	fmt.Printf("Type of the forth Slice is: %T\n", myFourthSlice)
	myFourthSlice = append(myFourthSlice, 1)
	fmt.Print(myFourthSlice, " ")
	fmt.Printf("Type of the forth Slice is: %T\n", myFourthSlice)
	myFourthSlice = append(myFourthSlice, 2, 3)
	fmt.Print(myFourthSlice, " ")
	fmt.Printf("Type of the forth Slice is: %T\n", myFourthSlice)
	b := myFourthSlice[1:3]
	fmt.Println(b)
	c := myFourthSlice[:3]
	fmt.Println(c)

	highScore := make([]int, 4)
	fmt.Println(highScore)

	highScore[0] = 2
	highScore[1] = 9
	highScore[2] = 4
	highScore[3] = 8
	// highScore[4] = 12
	fmt.Println("cap is: ", cap(highScore))
	highScore = append(highScore, 16, 17, 12, 18, 11, 21, 19)
	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Print("cap is: ", cap(highScore))
	fmt.Print(" len is: ", len(highScore))
	fmt.Println(" ", highScore)

	var pLangs = []string{"JS", "Python", "Java", "Ruby", "Rust", "Golang"}
	fmt.Println(pLangs)
	i := Find(pLangs, "Java")
	fmt.Println(i)

	// remove element based on index
	newpLang := append(pLangs[:i], pLangs[i+1:]...)
	fmt.Println(newpLang)

	var myXSlice = []int{1, 2, 3, 4, 5}
	myXSlice = myXSlice[1:3]
	fmt.Print(myXSlice, "\t")
	fmt.Print(cap(myXSlice), "\t\n")
	fmt.Println(myXSlice[:cap(myXSlice)])

}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}
