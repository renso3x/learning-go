package main

import (
	"fmt"
)

func main() {

	if firstRank, secondRank := 100, 50; firstRank > secondRank {
		fmt.Println("First is looking good")
	} else {
		fmt.Println("Second is looking good")
	}

	course := "go"

	switch course {
	case "js":
		fmt.Println("You are learning JS")
	case "go":
		fmt.Println("You are learning Go")
	default:
		fmt.Println("See pluralsight for learning tools")
	}

}
