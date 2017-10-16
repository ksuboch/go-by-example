package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) { // by convention, errors are the last return values
	if arg == 42 { // and have type error
		// constricting new error with message
		return -1, errors.New("Can't work with 42")
	}
	// nil value in the error position indicates that there was no errors
	return arg + 3, nil
}

// it's possible to use custom types as errors by implementing
// the Error() method on them
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// we use &argError syntax to build a new struct supplying
		// values for the two fields arg and prob
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 falied", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 falied:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
