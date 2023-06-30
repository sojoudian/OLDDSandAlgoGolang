package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("2006-01-02 Monday 15:04:05"))
	createDate := time.Date(1989, time.July, 07, 15, 04, 05, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("2006-01-02 Monday 15:04:05"))
}
