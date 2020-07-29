package main

import (
	"fmt"
	"math"
)

//package level constant
const counter int8 = 1
const (
	title             string  = "Hello There"
	flag              bool    = true
	packageLevelConst float32 = 3.14
)
const enumeratedCon = iota //iota use to create enumerated constant

//compiler infer automatically based on first iota
const (
	m1 = iota //0
	m2 = iota //1
	m3        //2
)

//iota reset at block level
const (
	m4 = iota //0
)

const (
	_         = iota             //0
	kiloBytes = 1 << (10 * iota) // 1 * 2^(10*1) = 1KB
	megaBytes                    // 1 * 2^(10*2) = 1MB
	gigaBytes
)

//const ExportedConst int = 25 //notice that first letter is uppercase hence it will be exported

// Can not use below declaration style at package level or at function levels for constant
// notUsed := true

func main() {
	fmt.Println(m1, m2, m3, m4, enumeratedCon)
	const counter int = 50 //re-declaration and shadowing

	//constant declaration
	const pi float64 = 3.14 //unused constant
	const a = 50            //using compiler ability to identify type
	var b int16 = 10
	const c int = 50

	var radius float64 = math.Sin(1)
	area := pi * radius * radius //expression is evaluated at compile time (compile time -> 3.14 * radius * radius)

	/* uncomment below lines to encounter error
	1. we can not redeclare constant under same scope
	2. We must define and initialize constant with some value at the same time
	3. We can not set value of constant whose value is determined at run time. Since constant are evaluated at compile time
	4. can not mix two types under same family unless identify using compiler ability. (int8 + int16 results in error)
	*/

	/*
		const counter int = 100

		const x string
		x = "Hey there"

		const y = math.Sin(60)
		fmt.Printf("y -> %v, x -> %v", y, x)

		fmt.Println("b + c -> ", c+b)
	*/

	fmt.Println("********************************")

	fmt.Printf("%v , Counter value -> %v , Counter type -> %T \n", title, counter, counter)
	fmt.Println("Area of circle -> ", area)
	fmt.Println("a + b -> ", a+b) //since compiler identify type for a, it puts value of "a" at compile time only. So no errors
	fmt.Printf("m1 -> %v, m2 -> %v, m3 -> %v, m4 -> %v, enumeratedCon -> %v \n", m1, m2, m3, m4, enumeratedCon)
	fmt.Println("*********************************")

}
