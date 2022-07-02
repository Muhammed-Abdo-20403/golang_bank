package main

import "fmt"

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {

	if knightIsAwake == true {
		panic(false)
	} else {
		fmt.Println("The knight is sleeping!")
	}
	return CanFastAttack(knightIsAwake)
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {

	if knightIsAwake == false || archerIsAwake == true && prisonerIsAwake == false {
		panic(true)
	} else {
		fmt.Println("characters is awake")
	}
	return CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake)
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {

	if archerIsAwake == false && prisonerIsAwake == true {
		panic(true)
	} else {
		fmt.Println("Can Signal Prisoner")
	}
	return CanSignalPrisoner(archerIsAwake, prisonerIsAwake)
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {

	if knightIsAwake == false && prisonerIsAwake == true || archerIsAwake == false && petDogIsPresent == true {
		panic(false)
	} else {
		fmt.Println("Can Free Prisoner")
	}
	return CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent)
}
