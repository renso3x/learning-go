package main

import (
	"fmt"
)

func main() {
	name := "Romeo"
	course := "Deep Dive GO"

	fmt.Println("Hi", name, "you are watching course ", course)

	changeCourse(&course)

	fmt.Println(course)

}

func changeCourse(course *string) string {
	*course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", course)
	return *course
}
