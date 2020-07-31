package main

import "fmt"

func main() {
	// general for loop
	for i := 0; i < 5; i++ {
		fmt.Println("index ", i)
	}

	// for loop with range to iterate over collection types
	stateMap := make(map[string]int)
	stateMap = map[string]int{
		"Mumbai":    1,
		"Chennai":   2,
		"Bangalore": 3,
	}
	for key, value := range stateMap {
		fmt.Printf("Map:: Key -> %v, Value -> %v \n", key, value)
	}

	stateSlice := make([]int, 0)
	stateSlice = append(stateSlice, 2, 3, 4)
	for key, value := range stateSlice {
		fmt.Printf("Slices:: Key -> %v, Value -> %v \n", key, value)
	}

	// for loop as a while loop
	j := 0
	for j < 2 {
		fmt.Println("for as a while -> index", j)
		j++
	}

	// multiple variable initialization
	for i, j := 0, 1; i < 2; i++ {
		fmt.Println("index i", i)
		fmt.Println("index j", j)
		j--
	}

	/* Uncomment below lines to encounter error
	1. we must put braces after for loop
	2. Invalid initialization format since we dont have , in loop in GO
	3. In Go, we can not put "," in expression since increment decrement is statement and not a expression
	*/

	/*
		for i := 0; i < 2; i++
			fmt.Println("Without braces ",i)

		for i := 0,j := 0, 1; i < 2; i++ {
			fmt.Println("index i", i)
			fmt.Println("index j", j)
			j--
		}

		for i,j:= 0,1; i < 2; i++,j-- {
			fmt.Println("index i", i)
			fmt.Println("index j", j)
		}
	*/

}
