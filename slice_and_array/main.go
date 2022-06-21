package main

import "fmt"

// import (
// "fmt"
// )
func main() {

	//// ********************************** ////
	//// ***** Using Slice and Arrays ***** ////
	//// ********************************** ////
	// A slice is declared just like an array,
	// but it doesn’t contain the size of the slice.
	// So it can grow or shrink according to the requirement.

	// Syntax
	// []T
	// or
	// []T{}
	// or
	// []T{value1, value2, value3, ...value n}

	// var my_slice[]int

	// Empty Silce
	keywords := []string{}
	fmt.Println(keywords)

	// Silce with string values
	names := []string{"Hassan", "4apo", "Yousef", "Moahmmed", "Ahmed", "Omer"}
	fmt.Println(names)

	// Update 0 index value
	names[0] = "Ibrahim"
	fmt.Println(names)

	// Extend slice
	names = append(names, "Addullah")
	fmt.Println(names)

	// Silce with Integer values
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(num)

	var num1 = num[0:9]
	num2 := num[0:]
	num3 := num[:6]
	num4 := num[:]
	num5 := num3[2:4]

	fmt.Println("Original Slice:", num)
	fmt.Println("New Slice 1:", num1)
	fmt.Println("New Slice 2:", num2)
	fmt.Println("New Slice 3:", num3)
	fmt.Println("New Slice 4:", num4)
	fmt.Println("New Slice 5:", num5)

	//// ********************************** ////
	//// *********** Using Arrays ********* ////
	//// ********************************** ////

	// array is created using the var keyword of
	// a particular type with name, size, and elements.

	// Syntax
	// Var array_name[length]Type
	// or
	// var array_name[length]Typle{item1, item2, item3, ...itemN}

	var name [3]string

	name[0] = "Mo"
	name[1] = "Abdo"
	name[2] = "Elshiek"

	fmt.Println("Elements of Array:")
	fmt.Println("Element 1: ", name[0])
	fmt.Println("Element 2: ", name[1])
	fmt.Println("Element 3: ", name[2])

	// Using shorthand declaration
	// array_name:= [length]Type{item1, item2, item3,...itemN}

	ages := [4]int{21, 19, 15, 5}

	fmt.Println("Element 2: ", ages[0])
	fmt.Println("Element 2: ", ages[1])
	fmt.Println("Element 2: ", ages[2])
	fmt.Println("Element 2: ", ages[3])

	for i := 0; i < 3; i++ {
		fmt.Println(ages[i])
	}

}
