package main

import "fmt"

func main() {
	executeIf()
	executeSwitch()
}

// If statement related code
func executeIf() {
	fmt.Println("********** If ************")
	// ###### simple if ###########
	if true {
		fmt.Println("This statement will execute")
	}
	if false {
		fmt.Println("This statement wont execute")
	}

	// ###### if else ###########
	var stateMap = make(map[string]int)
	stateMap = map[string]int{
		"Chennai": 1,
	}

	if _, ok := stateMap["Mumbai"]; ok {
		fmt.Println("mumbai is present in stateMap")
	} else {
		fmt.Println("Mumbai is not present in stateMap")
	}

	// ###### if - else if - else ###########

	num1 := 5
	num2 := 7

	if num1 > num2 {
		fmt.Println("Number 1 is greater than Number 2")
	} else if num1 < num2 {
		fmt.Println("Number 1 is less than Number 2")
	} else {
		fmt.Println("Number 1 is equal to Number 2")
	}

	/* Uncomment below lines to encounter error
	1. we must put braces after if clause
	2. Non boolean type not acceptable in if
	*/

	/*
		if false
			fmt.Println("Compiler error")
		if 3 {
			fmt.Println("true")
		}
	*/

	fmt.Println("*****************************")
	fmt.Println()

}

// Switch statement related code
func executeSwitch() {
	fmt.Println("********** Switch ************")

	// ###### generic switch (switch with tag) ###########
	// implicit break
	switch 1 {
	case 1:
		fmt.Println("Case 1 executed")
	case 2:
		fmt.Println("Case 2 executed")
	default:
		fmt.Println("Default case executed")
	}

	// ###### switch with multi case of same statements ###########
	switch 3 {
	case 1, 2, 3, 4, 5:
		fmt.Println("Single digit no between 1 to 5")
	case 6, 7, 8, 9:
		fmt.Println("Single digit no between 6 to 9")
	default:
		fmt.Println("Number is double digit")
	}

	// ###### switch case with string,map and type sensitivity ###########
	var city = make(map[string]string)
	city = map[string]string{
		"Mumbai":  "MI",
		"Chennai": "CSK",
	}
	switch city["mumbai"] {
	case "Mumbai":
		fmt.Println("City is Mumbai")
	case "Chennai":
		fmt.Println("City is Chennai")
	default:
		fmt.Println("City other than Mumbai or Chennai")
	}

	// ###### switch case with different type (not lossy conversion) ###########
	switch 1 {
	case 1.0:
		fmt.Println("int -> float + fallthrough case 1")
		fallthrough
	case 2.0:
		fmt.Println("int -> float + fallthrough case 2")
	default:
		fmt.Println("Number is double digit")
	}

	// ###### switch case without tags (untaged switch) ###########
	// Can override cases unlike tag switch statements
	i := 7
	switch {
	case i <= 5:
		fmt.Println("Less than 5")
	case i <= 10:
		fmt.Println("Less than 10")
	default:
		fmt.Println("Greater than 10")
	}

	var interfaceType interface{} = 1
	switch interfaceType.(type) {
	case int:
		fmt.Println("Integer type")
		break //explicit break (non mandatory)
	case float64:
		fmt.Println("Float64 type")
	case bool:
		fmt.Println("Boolean type")
	default:
		fmt.Println("Default type")
	}

	/* Uncomment below lines to encounter error
	1. Can not have duplicate case - compilation error
	2. If tag data type and case data type is mismatched in case of lossy conversion- compilation error
	*/

	/*
			switch 3 {
			case 1, 2, 3, 4, 5, 0:
				fmt.Println("Single digit no between 1 to 5")
			case 6, 7, 8, 9, 0:
				fmt.Println("Single digit no between 6 to 9")
			default:
				fmt.Println("Number is double digit")
			};

			switch 1 {
			case "1":
				fmt.Println("case 1")
			case "2":
				fmt.Println("case 2")
			default:
				fmt.Println("Default case")
		}
	*/

	fmt.Println("***************************")
	fmt.Println()
}
