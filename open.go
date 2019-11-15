package main

import (
	"fmt"
	"os"
)

// Error handling
func main() {
	_, err := os.Open("./main.go")

	if err != nil {
		fmt.Println("Error returned was:", err)
	}
}
