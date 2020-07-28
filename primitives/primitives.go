package main

import "fmt"

func main() {
	callBoolean()
	callInteger()
	callFloat()
	callComplex()
	callString()
	callRune()
}

func callBoolean() {
	var flag bool //Zero initialized value for boolean is false
	var trueFlag bool = true
	var falseFlag bool = 10 == 30 //output of comparison is boolean

	// uncomment below lines to encounter error (boolean can not used as alias for other type)
	/*
		var i bool = -1
		var j bool = bool(-1)
		fmt.Println("Value of i -> %v , j -> %v", i, j)
	*/

	fmt.Println("********** Boolean ************")

	fmt.Printf("Value of flag is %v \n", flag)
	fmt.Printf("Type of flag is %T \n", trueFlag)
	fmt.Println("Comparison Value ->", falseFlag)

	fmt.Println("*********************************")
	fmt.Println()

}

func callInteger() {
	var i int //Zero initialized value + min 32 bits occupied
	var j uint16 = 15510
	var x int = 10
	var y = 20 //allowed. bydefault int will be taken. if decimal points present then float will be taken

	var a uint = 10
	var b uint = 3

	// uncomment below lines to encounter error (can not mix two types under same family)
	/*
		var m int = 50
		var n int16 = 30
		fmt.Println("m + n ->", m+n)
	*/

	fmt.Println("********** Integer ************")

	fmt.Printf("Value of i is %v \n", i)
	fmt.Printf("Type of j is %T and binary is %b \n", j, j)
	fmt.Println("X + Y ->", x+y)
	fmt.Println("a / b ->", a/b)
	fmt.Println("a % b ->", a%b)

	fmt.Println("*********************************")
	fmt.Println()

}

func callFloat() {
	var i float32 //Zero initialized value + min 32 bits occupied
	j := 1.5e4    // default type is float64
	var x float32 = 10
	var y float32 = 20.2
	var u = 3.14 //allowed. default float64 will be taken

	// uncomment below lines to encounter error (can not mix two types under same family)
	/*
		var m float32 = 50.5
		var n float64 = 30.6
		l := 30.0
		fmt.Println("m + n ->", m+n)
		fmt.Println("m + l ->", m+l)
		fmt.Println("l + n ->", l+n) // this will work because ":=" will have float64 by default
	*/

	fmt.Println("********** Float32,64 ************")

	fmt.Printf("Value of i is %v \n", i)
	fmt.Printf("Value of u is %v and type is %T\n", u, u)
	fmt.Printf("Type of j is %T and binary is %b \n", j, j)
	fmt.Println("X + Y ->", x+y)
	fmt.Println("y / x ->", y/x)

	fmt.Println("*********************************")
	fmt.Println()
}

func callComplex() {
	var i complex64 //Zero initialized value + min 64 bits (32 for real + 32 for imaginary parts) occupied
	j := 80 + 80i   //default is complex128
	var x complex64 = 10 + 10i
	var y complex64 = 20.2 + 20i
	var u complex128 = complex(5, 50) //using complex function

	// uncomment below lines to encounter error (can not mix two types under same family)
	/*
		var m complex64 = 50 // (50+0i)
		var n complex128 = 30
		fmt.Println("m + n ->", m+n)
	*/

	fmt.Println("********** Complex64, 128 ************")

	fmt.Printf("Value of i is %v \n", i)
	fmt.Printf("Type of j is %T and binary is %b \n", j, j)
	fmt.Println("X + Y ->", x+y)
	fmt.Println("y / x ->", y/x)
	fmt.Printf("Real part: type of u is %T and value is %v \n", real(u), real(u))
	fmt.Printf("Imaginary part: type of u is %T and value is %v \n", imag(u), imag(u))

	fmt.Println("*********************************")
	fmt.Println()
}

func callString() {
	var x string = "hello"
	y := []byte(x) // converting to array of bytes
	var z string = "There"

	// uncomment below lines to encounter error (strings are immutable)
	/*
		var m string = "Hi"
		m[1] = "e"
		fmt.Println("m ->", m)
	*/

	fmt.Println("********** String ************")

	fmt.Printf("value of x -> %v, type of x -> %T \n", x, x)
	fmt.Printf("value of character of x -> %v, type of character of x -> %T \n", x[1], x[1]) // string as a alias of bytes
	fmt.Printf("value of y -> %v, type of y -> %T \n", y, y)
	fmt.Println("X + Z ->", x+" "+z) // string concatination

	fmt.Println("*********************************")
	fmt.Println()

}

func callRune() {
	var letter rune = 'a' //Single quote for rune, double quotes for string

	// uncomment below lines to encounter error (only single character is allowed)
	/*
		var m rune = 'Hi'
		fmt.Println("m ->", m)
	*/

	fmt.Println("********** Rune ************")

	fmt.Printf("value of letter -> %v, type of letter -> %T \n", letter, letter) //rune are type alias of int32

	fmt.Println("*********************************")
	fmt.Println()
}
