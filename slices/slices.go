package main

import "fmt"

func main() {
	var slice []int = []int{1, 2, 3, 4, 5, 6} // []int on left-side is not mandatory
	copyArr := slice                          //copyArr is pointing to same underlying arr (address is same)
	copyArr[2] = 50

	referenceArr := &slice //referenceArr is pointing to same arr (address is same)

	var length int = 3
	var capacity int = 10
	var makeSlice = make([]int, length, capacity) //initially length no of elements with zero value ([0 0 0])
	makeSlice = append(makeSlice, 2, 3, 4)        //[0 0 0 2 3 4], len = 6 , cap = 10
	// We can use spread operator kind which will spread all element of slices
	makeSlice = append(makeSlice, []int{7, 8, 9, 10, 11}...) //len = 11 , cap = 20

	var animalArr []string
	animalArr = append(animalArr, "cat")

	/* uncomment below lines to encounter error
	1. const can not be used with slice
	2. We can not directly use slices under slices for append
	3. Runtime error will occur if we define slice and then set values at index without append(unlike arrays) because length will be considered 0
	*/

	/*
		const m = []int{1,2,3}
		makeSlice = append(makeSlice, []int{7, 8, 9})
		var animalArr []string
		animalArr[0] = "cat"

		fmt.Println("Slice with const - size declared and initialized ", m)
		fmt.Println("Slice with make function and append method with slice as a args ",makeSlice)
	*/

	fmt.Println("********************************")

	fmt.Println("Array with size declared and initialized ", slice, "Length ", len(slice), " Capacity ", cap(slice))
	fmt.Println("Copy array ", copyArr)
	fmt.Println("reference array ", *referenceArr) //derefenced with *
	fmt.Println("Array with append func", animalArr)
	fmt.Println("Array with make function ", makeSlice, "Length ", len(makeSlice), " Capacity ", cap(makeSlice))
	callSlice(slice)

	fmt.Println("********************************")
	fmt.Println()
}

func callSlice(arr []int) {
	b := arr[:]   //slices of all element
	c := arr[3:]  // slices of element starting from 4th element since array starts with 0
	d := arr[:5]  // slices of element starting from first element to 5th element (0-4)
	e := arr[3:5] // slice of element from 3rd index(inclusive) to 5th index(exclusive)
	fmt.Println("Array under callSlice -> ", arr)
	fmt.Printf("Slices under callSlice: %v, %v, %v, %v \n", b, c, d, e)
	e[1] = 90 //Changing value at index 1 of e which is index 4 in backed array, hence this will change at everywhere at index 4
	fmt.Println("After value change in one of the slice: array under callSlice -> ", arr)
	fmt.Printf("After: value change in one of the slice: slices under callSlice: %v, %v, %v, %v \n", b, c, d, e)
}
