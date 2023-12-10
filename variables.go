package main

import "fmt"

func main() {

	// Defining variables
	// There are different ways to define variables

	// 1. var <variable_name> <variable_type> = <value>
	var name string = "John"

	// 2. var <variable_name> = <value>
	var name2 = "Esi"

	// 3. <variable_name> := <value>
	// NOTE: Doesn't work outside a function
	name3 := "Ama"

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)

}
