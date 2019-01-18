package main

import "fmt"

func main() {
	// define map
	emails := make(map[string]string)

	emails["Bob"] = "bob@abc.com"
	emails["Alice"] = "alice@abc.com"
	emails["Mike"] = "mike@abc.com"

	fmt.Println(emails)

	// delete
	delete(emails, "Bob")
	fmt.Println(emails)

	mymap := map[int]string{1: "A", 2: "B"}
	fmt.Println(mymap)
}
