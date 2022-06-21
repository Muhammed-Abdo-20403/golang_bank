package main

import "fmt"

func main() {

	//// ******************************************* ////
	//// ************** Using Maps ***************** ////
	//// ******************************************* ////

	// Maps is a collection of unordered pairs of key-value.
	// It is widely used because it provides fast lookups and values that can retrieve,
	// update, or delete with the help of keys.

	// Syntax:
	// An Empty map
	// map[Key_Type]Value_Type{}

	// Map with key-value pair
	// map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

	animals := map[int]string{
		1: "Dog",
		2: "Cad",
		3: "Caw",
		4: "Goat",
		5: "Ant",
		6: "Finch",
		7: "Canary",
		8: "Horse",
	}

	fmt.Println(animals)

	animals[9] = "Monkey"

	fmt.Println(animals)

	// delete a value
	delete(animals, 9)
	fmt.Println(animals)

	// delete a fack value
	delete(animals, 12)
	fmt.Println(animals)

	// create a map
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	// insert key-value pairs in the map
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	fmt.Println(countryCapitalMap)

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

}
