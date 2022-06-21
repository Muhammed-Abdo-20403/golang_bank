/* My first program in go */
package main

import "fmt"

func main() {

	//// *********************************************** ////
	////*************** Usering For loop *************** ////

	//// ******* Syntax ******* ////
	// for initialization; condition; post{
	//     // statements....
	// }

	//// ******* For loop as Infinite Loop ******* ////
	// for {
	// 	// Statement...
	// }

	//// ******* Infinite Loop ******* ////

	// for {
	// 	fmt.Printf("Hello Mo Salah!")
	// }

	colors := []string{"Navy", "Lime", "Purple", "Blue", "Off White", "Gray", "Black", "White", "Green", "Yellow", "Red"}

	for countColor, perColor := range colors {
		fmt.Println(countColor, ":", "It's", perColor)
	}

	for num := 0; num <= 10; num++ {
		for name := 0; name <= 5; name++ {
			fmt.Println(num, "for", name)
		}
	}

	for employee := 0; employee <= 5; employee++ {
		if employee == 4 {
			fmt.Println("You Are Num 4")
			break
		}
	}

}
