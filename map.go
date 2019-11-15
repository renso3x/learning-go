package main

import "fmt"

func main() {

	// map standings
	nbaTeams := map[string]int{
		"Lakers":  1,
		"Celtics": 2,
	}

	// for loop
	for i, v := range nbaTeams {
		fmt.Println("\nKey: %i, Value: %v", i, v)
	}

	nbaTeams["Lakers"] = 10

	fmt.Println(nbaTeams["Lakers"])

	delete(nbaTeams, "Lakers")

	fmt.Println(nbaTeams["Lakers"])
	fmt.Println(nbaTeams)

	append(nbaTeams, {"Portland": 3})

	fmt.Println(nbaTeams)
}
