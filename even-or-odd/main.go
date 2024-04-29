package main

import "fmt"

func main() {
	sliceSize := 11
	mySlice := []int{}

	for i := 0; i < sliceSize; i++ {
		mySlice = append(mySlice, i)
	}

	for _, number := range mySlice {
		if number % 2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}