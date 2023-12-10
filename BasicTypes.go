package main

import "fmt"

func basicTypes() {
	// Examples of the types in Go
	// int has several sizes like int8, int16, int32, int64
	var v1 int = 10

	// float also has several sizes like float32, float64
	var v2 float64 = 10.5

	// string
	var v3 string = "Hello"

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
}
