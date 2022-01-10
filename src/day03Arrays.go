package main

import "fmt"

func arrays() {

	scores := [4]int{1, 3, 5, 7}
	for index, value := range scores {
		fmt.Printf("sorce[%d] is %d\n", index, value)
	}
}

func main() {
	arrays()
}
