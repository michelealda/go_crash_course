package main

import "fmt"

func main() {
	ids := []int{11, 22, 33, 44, 55, 66}

	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	for id := range ids {
		fmt.Printf("ID %d\n", id)
	}

	// Sum ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//using maps
	mymap := map[int]string{1: "A", 2: "B"}
	for k, v := range mymap {
		fmt.Printf("%d: %s\n", k, v)
	}

}
