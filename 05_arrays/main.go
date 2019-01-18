package main

import "fmt"

func main() {
	//Arrays
	var fruitArray [2]string

	fruitArray[0] = "apple"
	fruitArray[1] = "pear"

	//Declare and assign
	shapeArray := [2]string{"square", "circle"}

	fmt.Println(fruitArray)
	fmt.Println(shapeArray)

	//Slices

	fruitSlice := []string{"apple", "pear", "orange"}
	fmt.Println(fruitSlice)

}
