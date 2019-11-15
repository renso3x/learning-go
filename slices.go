package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(mySlice[4])

	mySlice[1] = 0 //change value in index of 1
	fmt.Println(mySlice)

	fmt.Println(mySlice[2:5])            // slice of slice
	fmt.Println(mySlice[len(mySlice)-1]) // end of the array

	for _, i := range mySlice {
		fmt.Println(i)
	}

	// appended new slice
	newSlice := []int{4, 6, 3}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)
}
