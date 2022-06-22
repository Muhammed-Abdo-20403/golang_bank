package main

import "fmt"

// Function doesn't returne anytihing
func names(name string) {
	fmt.Println(name)
}

// Function take tow int args and return the sum
func plus(num1, num2 int) int {
	return num1 + num2 // bc return type
}

// Same above
func plusPlus(num1, num2, num3 int) int {
	return num1 + num2 - num3
}

func minus(params ...interface{}) {
	fmt.Println(len(params))
}

// Muliple results will return
func multiResult(ft, sec string) (string, int) {
	res := ft + sec
	return res, len(res)
}

// Same above
func multiResult2(ft, sec string) (res string, l int) {
	res = ft + sec
	l = len(res)
	return
}

func nums(frt, secn int) int {

	var thd int
	thd = frt
	frt = secn
	secn = thd

	return thd
}

func main() {
	///// ******************************************** ////
	//// ************** Using Functions ************** ////

	// function is a collection of statements that perform
	// some specific task and return the result to the caller

	// Syntax
	// func function_name(Parameter-list)(Return_type){
	// 	function body.....
	// }

	names("Mo Elshiekh")

	num := plus(5, 6)
	fmt.Println(num)

	num = plusPlus(3, 5, 6)
	fmt.Println(num)

	minus("mo", 22, 2.4, true)

	ft1, sec2 := multiResult("Hassan", "Sayed")
	fmt.Println(ft1, sec2)

	ft1, sec2 = multiResult2("Hassan", "Sayed")
	fmt.Println(ft1, sec2)

	var fth int = 10
	var fith int = 20
	fmt.Println("fth =", fth, "and", "fith = ", fith)

	nums(fth, fith)
	fmt.Println("fth =", fth, "and", "fith = ", fith)
}
