package main

import (
	"errors"
	"fmt"
)

// walker type define something that can walk
// type walker interface {
// 	walk(p point) error
// 	getPosition() point
// }

// talker type define something that can talk
// type talker interface {
// 	talk(s string)
// }

// type point struct {
// 	x int
// 	y int
// }

// human struct type
// type human struct {
// 	position point
// }

// implement interface walker on human
// interface implementation comes implicitly when struct implement all the functions in interface type
// func (h *human) walk(p point) error {
// 	if p.x < 0 || p.y < 0 {
// 		return errors.New("invalid point")
// 	}
// 	h.position = p
// 	fmt.Println("Human walked to", h.position)
// 	return nil
// }
// func (h *human) getPosition() point {
// 	return h.position
// }

// implement interface talker on human
// func (h *human) talk(s string) {
// 	fmt.Println("Human talking", s)
// }

// type animal struct {
// 	position point
// }

// func (a *animal) walk(p point) error {
// 	if p.x < 0 || p.y < 0 {
// 		return errors.New("invalid point")
// 	}
// 	a.position = p
// 	fmt.Println("animal walked to", a.position)
// 	return nil
// }
// func (a *animal) getPosition() point {
// 	return a.position
// }

// func main() {
// 	eslam := &human{}
// 	dog := &animal{}

// 	steps := []point{
// 		point{1, 1},
// 		point{2, 2},
// 		point{3, 3},
// 	}
// 	err := move(eslam, steps)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = move(dog, steps)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = move(dog, []point{point{-1, 2}})
// 	if err != nil {
// 		fmt.Println(err)
// }

// Functions should take interface as parameter type
// that makes it work better with diferent type of sturcts
// this will work
// moveHuman(eslam, steps)
// this will cause errors
// cannot use dog (type *animal) as type *human in argument to moveHuman
// moveHuman(dog, steps)
// }

// function takes interface as parameter type
// can take any struct type implements the interface
// func move(w walker, points []point) error {
// 	for _, p := range points {
// 		err := w.walk(p)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// function takes struct as type for the parameter
// can not take any other struct even if it implements the required functions
// func moveHuman(w *human, points []point) error {
// 	for _, p := range points {
// 		err := w.walk(p)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
func getError(b bool) (err error) {
	if b == true {
		// err = errors.New("this is an error :(")
		err = fmt.Errorf("this is an error :(")
	}
	return
}

type invalidInput struct {
	filedName string
}

func (i invalidInput) Error() string {
	return fmt.Sprintf("invlaid input on field '%v'", i.filedName)
}

func validateInput(field, value string) (err error) {
	if value == "" {
		err = invalidInput{filedName: field}
	}
	return
}

func main() {
	err := errors.New("It Says Error!")
	if err != nil {
		fmt.Println(err)
	}

	err = errors.New("Sample Error")
	if err != nil {
		fmt.Println(err)
	}

	{
		// this is how we check on error
		err := getError(true)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	{
		// call to validateInput function with valid input
		err := validateInput("firstName", "eslam")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	{
		// call to validateInput function with invalid input
		err := validateInput("lastName", "")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	{
		// check on error value
		// err := invalidLength("")
		err := invalidLength("error")
		// err := getError(true)
		if err == errInvalidLength {
			fmt.Println("We know err value using 'err == errInvalidLength'")
		} else {
			fmt.Println("we do not know about this error")
		}
	}

	{
		// check on error type using "type assertion"
		// this works only with custom types like invalidInput
		// does not work with variable like errInvalidLength
		err := validateInput("name", "")
		_, ok := err.(invalidInput)
		if ok {
			fmt.Println("We know err type using 'type assertion'")
		} else {
			fmt.Println("we do not know about this error")
		}
	}
}

var errInvalidLength = errors.New("Invalid length")

func invalidLength(s string) error {
	if len(s) < 1 {
		return errInvalidLength
	} else if s == "error" {
		return errors.New("Invalid input")
	}
	return nil
}
