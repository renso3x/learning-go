package main

import (
	"fmt"
	"time"
)

func main() {
	//self destruct a loop
	for i := 10; i >= 0; i-- {

		if i%2 == 0 {
			// fmt.Println("Boom!")
			continue
		}

		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	// rangeLoop()

}

// match the coures with the courseInProgress
func rangeLoop() {
	courses := []string{"Docker Deep Dive", "Docker Fundamentals", "Docker: Go"}

	courseInProgress := []string{"Python 3: Bootcamp", "Docker Deep Dive", "Dockerize Everything"}

	for _, course := range courses {
		for _, j := range courseInProgress {
			if course == j {
				fmt.Println("Hey, we found something similar " + course)
				break
			}
		}
	}

}
