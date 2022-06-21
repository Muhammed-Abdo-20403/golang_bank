/* My first program in go */
package main

import "fmt"

func main() {

	//// ********************************************* ////
	//// ************* Using If condition ************ ////

	var soccerPlayersName1 string = "CR7"
	var soccerPlayersAge1 int = 35
	if soccerPlayersName1 == "CR7" && soccerPlayersAge1 == 35 {
		fmt.Println("Hi", soccerPlayersName1, "You Are", soccerPlayersAge1)
	} else {
		fmt.Println("Who are you?")
	}

	fmt.Println("**************************************")

	fmt.Println("Inter Your Name And Your Age!")
	soccerPlayersName := "Messi"
	soccerPlayersName2 := "Mo Salah"
	soccerPlayersAge, soccerPlayersAge2 := 34, 30

	if soccerPlayersName == "Messi" && soccerPlayersAge2 == 34 {
		fmt.Println("Welcome,", soccerPlayersName)
		fmt.Println("You Are", soccerPlayersAge, "Old")
	} else if soccerPlayersName2 == "Mo Salah" || soccerPlayersAge2 == 34 {
		fmt.Println("Welcome,", soccerPlayersName2)
		fmt.Println("You Are", soccerPlayersAge2, "Old")
	} else {
		fmt.Println("Not Exsist, Plss Getout")
	}
}
