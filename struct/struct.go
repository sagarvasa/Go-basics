package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	name  string
	id    int
	roles []string
}

type bank struct {
	name    string
	code    string
	pincode int
}

type stateBank struct {
	bank          //embedding bank struct into stateBank struct
	branchcode    int
	withdrawLimit float64
}

// struct fields with tags (anything between ``)
type animal struct {
	name  string ` required:"true" custom:"50" max:"50"`
	breed string
}

func main() {

	employee := employee{
		id:    123,
		name:  "Sagar",
		roles: []string{"technology", "product"},
	}

	/* We can skip keys in struct initialization but its not recommended as
	we need to add value in positional manner and if new value is added in struct then we need to maintain same order everywhere
	*/

	/*
		employee2 := employee{
			"Monu",
			123,
			[]string{"accounting", "finance"},
		}

		fmt.Println("Struct by Position ", employee2)
	*/

	/* uncomment below lines to encounter error
	1. Position mismatch since key is not present
	2. Too less values passed in struct in case of absent of key
	*/
	/*
		employee3 := employee{
			123,
			"Monu",
			[]string{"accounting", "finance"},
		}
		employee4 := employee{
			"Sagar",
			123,
		}
		fmt.Println("Struct by Position ", employee3)
		fmt.Println("Struct with less parameter ", employee4)
	*/

	//anonymous struct
	anonymousStruct := struct{ name string }{name: "Sagar"}

	copyStruct := anonymousStruct //struct are value type
	copyStruct.name = "Sam"

	referenceStruct := &anonymousStruct //now pointing to same memory as of anonymousStruct

	bankStruct := stateBank{}
	bankStruct.name = "State Bank of India" //belongs to bank struct
	bankStruct.code = "SBI"                 //belongs to bank struct
	bankStruct.pincode = 123456             //belongs to bank struct
	bankStruct.branchcode = 654321
	bankStruct.withdrawLimit = 9999999.99

	// alternate to above bank struct
	//bankStruct.bank = bank{name: "State Bank of India", code: "SBI", pincode: 123456}

	// using reflect library to get tags under struct
	st := reflect.TypeOf(animal{})
	fmt.Println(st)
	animalName, ok := st.FieldByName("name")

	//animalName, ok := st.FieldByName("wrongFieldName") //animalName = "", ok = false

	fmt.Println("********************************")

	fmt.Println(employee)
	fmt.Println("anonymousStruct ", anonymousStruct)
	fmt.Println("copyStruct ", copyStruct)
	referenceStruct.name = "Monu"
	fmt.Println("after name change in reference struct : anonymousStruct ", anonymousStruct)
	fmt.Println("referenceStruct ", *referenceStruct)
	fmt.Println("Embedded struct ", bankStruct)

	fmt.Println("Tags ", animalName.Tag, "\nisFound ", ok)

	fmt.Println("********************************")
	fmt.Println()
}
