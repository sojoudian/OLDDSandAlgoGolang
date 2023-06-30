package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please inter your message: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your message, ", input)
	fmt.Printf("Type of the input is %T \n", input)
}
