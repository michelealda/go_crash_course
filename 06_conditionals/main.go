package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x <= y {
		fmt.Printf("%d is less than %d", x, y)
	}

	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	}

}
