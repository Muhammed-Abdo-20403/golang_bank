package main

import "fmt"

// Declaring a structure:
// type Address struct {
//       name string
//       street string
//       city string
//       state string
//       Pincode int
// }

// Or

// type Address struct {
//     name, street, city, state string
//     Pincode int
// }

type Player struct {
	ftName  string
	lstName string
	score   int
}

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {
	//// *************************************** ////
	//// ********** Using Structures *********** ////
	//// *************************************** ////

	// structure or struct is a user-defined type that allows
	// to group/combine items of possibly different types into a single type.

	score1 := Player{"Mo", "Salah", 5}
	fmt.Println(score1)

	ftscore := &score1

	fmt.Println(ftscore)

	fmt.Println(ftscore.ftName)
	fmt.Println(ftscore.lstName)
	fmt.Println(ftscore.score)

	ftscore.ftName = "CR7"
	ftscore.lstName = "Ronaldo"
	ftscore.score = 7

	fmt.Println(ftscore.ftName)
	fmt.Println(ftscore.lstName)
	fmt.Println(ftscore.score)

	var a Address
	fmt.Println(a)

	adrss1 := Address{"Ahmed Elmenofy", "Elsalam City", 3623572}
	fmt.Println("Address1: ", adrss1)

	adrss2 := Address{"Elmoassas", "Elmaj City", 21545488}
	fmt.Println("Address2: ", adrss2)

}
