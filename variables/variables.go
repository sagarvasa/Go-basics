package main

import "fmt"

//package level variables
var counter int = 1
var (
	title           string  = "Hello There"
	flag            bool    = true
	packageLevelVar float32 = 3.14
)

//var ExportedVar int = 25 //notice that first letter is uppercase hence it will be exported

// Can not use below variable declaration style at package level
// notUsed := true

func main() {
	var counter int = 50 //re-declaration and shadowing

	//variable declaration
	var i int = 5
	var j float32 //Zero value initialization
	k := 10
	var x string

	x = "Hey there"

	/* uncomment below lines to encounter error
	1. we can not redeclare variable under same scope
	2. All variables must be used to avoid compilation errors
	*/

	/*
		var counter int = 100
		var unusedFlag bool
	*/

	fmt.Println("********************************")

	fmt.Printf("%v, I am %T \n", x, x)
	fmt.Println(title+", Counter value -> ", counter)
	fmt.Printf("i -> %v, j -> %v, k -> %v \n", i, j, k)

	checkVisibility()

	fmt.Println("*********************************")

}

func checkVisibility() {
	var blockVar float32 = 90

	fmt.Println("Package level variable -> ", packageLevelVar)
	fmt.Println("Block Level variable -> ", blockVar)
	//fmt.Println("Exported variable -> ", ExportedVar)
}
