package main

import "fmt"

func compositeTypes() {

	// Defining an array
	// Arrays can be of any basic type or composite type int, string, float, struct, array, slice, map, interface
	var arr = [5]int{1, 2, 3, 4, 5}

	// Defining a slice
	// Slices can be of any basic type or composite type int, string, float, struct, array, slice, map, interface
	var slice = []int{1, 2, 3, 4, 5}

	// Defining a map
	// Maps can be of any basic type or composite type int, string, float, struct, array, slice, map, interface
	var map1 = map[string]int{"a": 1, "b": 2, "c": 3}

	// Defining a struct
	// Structs can be of any basic type or composite type int, string, float, struct, array, slice, map, interface
	// A struct is a collection of fields
	type person struct {
		name string
		age  int
	}

	type employee struct {
		name string
		age  int
		id   int
		pay  float32
	}

	// Using a struct
	// Using the person struct
	Kofi := person{"Kofi", 20}

	// Using the employee struct
	Employee1 := employee{"John", 20, 1234, 1000.00}

	fmt.Println("Array: ", arr)
	fmt.Println("Slice: ", slice)
	fmt.Println("Map", map1)
	fmt.Println("Struct: ", Kofi)
	fmt.Println("Struct: ", Employee1)

	// Interfaces are a bit complex, I'm going to skip it for now
}
