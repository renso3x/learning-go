package main

import "fmt"

func main() {

	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	GoFunds := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Beginner",
		Rating: 3.4,
	}

	fmt.Println(GoFunds.Author)
}
