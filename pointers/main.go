package main

import "fmt"

func main() {
	//// ***************************************** ////
	//// *********** Using Poinater ************** ////
	//// ***************************************** ////

	// Pointer is variable that is used to store
	// the memory address of another variable.

	// * Operator also termed as the dereferencing operator used
	// to declare pointer variable and access the value stored in the address.

	// & operator termed as address operator used to returns the address
	// of a variable or to access the address of a variable to a pointer.

	i, j := 50, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// taking a normal variable
	var x int = 5748

	// declaration of pointer
	var t *int

	// initialization of pointer
	t = &x

	// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", t)

	// using := operator to declare
	// and initialize the variable
	y := 458

	// taking a pointer variable using
	// := by assigning it with the
	// address of variable y
	c := &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", c)

}
