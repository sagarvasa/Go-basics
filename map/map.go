package main

import "fmt"

func main() {
	var directMap map[string]int = map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	} //map[string]int on left side is non mandatory
	copyMap := directMap //copyMap is pointing to same underlying map (address is same)
	copyMap["d"] = 100

	referenceMap := &directMap //referenceMap is pointing to same map (address is same)

	var stateCodeMap = make(map[string]string)
	stateCodeMap = map[string]string{
		"Mumbai":    "MI",
		"Bangalore": "RCB",
		"Chennai":   "CSK",
		"Rajasthan": "RR",
		"pune":      "RPS",
	}
	delete(stateCodeMap, "pune") // delete key from map
	stateCodeMap["Pune"] = "RPS" //add key to map

	var mumbaiCity = stateCodeMap["Mumbai"]
	var bangaloreCity = stateCodeMap["bangalore"] //returns zero value of data type if not found

	chennaiCity, ok := stateCodeMap["chennai"] // ok = false then value didnt found.
	if !ok {
		fmt.Println("Wrong city name")
	}

	// uncomment below lines to encounter error (const can not be used with map)
	/*
		const m = map[string]int{"a": 97, "b": 98}
		fmt.Println("Map with const - size declared and initialized ", m)
	*/

	fmt.Println("********************************")

	fmt.Println("map declared and initialized", directMap, "Length ", len(directMap))
	fmt.Println("Copy map", copyMap)
	fmt.Println("reference map", *referenceMap) //derefenced with *
	fmt.Println("Map with make function", stateCodeMap)
	fmt.Println("Map without ok variable: found case:", mumbaiCity)
	fmt.Println("Map without ok variable: not found case:", bangaloreCity)
	fmt.Println("Map with ok variable: not found case:", chennaiCity, ok)

	fmt.Println("********************************")
	fmt.Println()
}
