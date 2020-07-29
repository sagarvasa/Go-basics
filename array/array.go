package main

import "fmt"

func main() {
	var arr = [3]int{1, 2, 3}
	copyArr := arr //All values are copied from arr to copyArr
	copyArr[2] = 5

	referenceArr := &arr //referenceArr is pointing to same arr (address is same)

	var fruitArr = [...]string{"apple", "mango", "banana", "grapes"} //since size is already known, we can skip it with ...

	var animalArr [3]string
	animalArr[0] = "cat"
	animalArr[1] = "dog"
	animalArr[2] = "cow"

	// uncomment below lines to encounter error (const can not be used with array, unused variable)
	/*
		const m = [3]int{1,2,3}
		var unusedArr [3]int
		fmt.Println("Array with const - size declared and initialized ", m)
	*/

	fmt.Println("Array with size declared and initialized ", arr, "Length ", len(arr), " Capacity ", cap(arr))
	fmt.Println("Copy array ", copyArr)
	referenceArr[2] = 9
	fmt.Println("Original arr after change in reference arr ", arr)
	fmt.Println("reference array ", *referenceArr) //derefence with *
	fmt.Println("Array with ... as size and initialized ", fruitArr)
	fmt.Println("Array with size declared and initialized later ", animalArr)
}
