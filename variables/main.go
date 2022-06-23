/* My first program in go */
package main

import "fmt"

func main() {
	//// ************************************************* ////
	//// ************* Using Variables ******************* ////

	// Using and declarating variables
	var playerName string // string var
	var playerAge int     // intger var
	playerName = "Mo Salah"
	playerAge = 30
	clubName := "Liverbool Club"
	country := "Egypt"
	spouse := "Magi Sadeq"
	children := 2

	childrenName := "Makka Mohamed Salah and Kayan Mohamed Salah" // This var is string
	salary := 2.5                                                 // This var is float

	// Output in terminal
	fmt.Println("Hello,", playerName)
	fmt.Println("You", playerAge, "years old.")
	fmt.Println("You're From", country)
	fmt.Println("Your spouse is", spouse)
	fmt.Println("You have", children, "and they are", childrenName)
	fmt.Println("You Play for", clubName)
	fmt.Println("Your salary is", salary, "Milion")

	fmt.Println("**************************************")

	name1 := "Mo Elsheikh"
	name2 := "Mo Elsheikh"
	name3 := "Mo Elsheikh"
	result1 := name1 == name2
	result2 := name1 == name3

	// Display the result
	fmt.Println(result1)
	fmt.Println(result2)

	// Display
	fmt.Println("Length of the string is: ", len(name1))

	// Display the type of result1 and result2
	fmt.Println("The type of result1 is and "+"the type of result2 is ", result1, result2)

	fmt.Println("**************************************")

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)
	// alternative syntax
	i = 333
	f = float64(i)
	u = uint(f)

	fmt.Println(i, f, u)

	fmt.Println("**************************************")

	//// *************************************** ////
	//// ******** Using Variable Scope ********* ////
	//// *************************************** ////

	// Local Variables(Declared Inside a block or a function)
	// Global Variables(Declared outside a block or a function)

	num := 5
	fmt.Println("Num Before", num)
	{
		// available accessible
		num = 5
		fmt.Println("Num In:", num)
	}
	fmt.Println("Num Out After:", num)

	num1 := 20
	fmt.Println("New Num out", num1)
	{
		// Declared a new variable with same nam
		num1 := 5
		fmt.Println("New Num In:", num1)

		// Update in Variable
		num1 = 15
		fmt.Println("Update Num In:", num1)
	}

	// update out variable
	num1 = 30
	fmt.Println("Update Num Out:", num1)

	// will cause error "undfefined: mo"
	// fmt.Print("Mo Salah", mo)
	// mo := "Mo Salah"

}
