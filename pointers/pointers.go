package main

import "fmt"

func main() {
	fmt.Println("************* Pointers **************")
	callPointer()
	fmt.Println("*************************************")
}

type customStruct struct {
	name   string
	gender string
}

func callPointer() {
	var i int = 1
	var j *int = &i // *int is pointer to int; &a is address of a
	var nilPointer *int

	var str *customStruct
	str = new(customStruct)
	//str = &customStruct{name: "Sagar", gender: "Male"} //alternative to above one
	(*str).name = "Sagar"
	str.gender = "Male" //same as above. compiler automatically infer as above

	fmt.Printf("Value of i -> %v, j -> %v, type of j -> %T\n", i, j, j)
	fmt.Println("Value of j after dereferencing ", *j) //*j is derefering j
	fmt.Println("Value of pointer without assigning ", nilPointer)
	fmt.Println("Struct with new initializer", *str)

	/* Uncomment below lines to encounter error
	1. Pointer arithmatic not allowed in GO (int* and int is type mismatch)
	*/

	/*
		a := 5;
		b := &a + 5 //-> Error
	*/

}
