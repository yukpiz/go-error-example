package main

import (
	"fmt"
	"log"

	"golang.org/x/xerrors"
)

func main() {
	err := ErrorFn()
	log.Printf("%v\n", err)
	log.Printf("%+v\n", err)
	fmt.Printf("%v\n", err)
	fmt.Printf("%+v\n", err)
}

func ErrorFn() error {
	return ErrorCodes.ErrorType1.Error
}

type customError struct {
	Error       error
	Description string
}

type errs struct {
	ErrorType1 customError
	ErrorType2 customError
	ErrorType3 customError
	ErrorType4 customError
	ErrorType5 customError
}

var ErrorCodes = errs{
	ErrorType1: customError{Error: xerrors.New("Custom.Error1")},
	ErrorType2: customError{Error: xerrors.New("Custom.Error2")},
	ErrorType3: customError{Error: xerrors.New("Custom.Error3")},
	ErrorType4: customError{Error: xerrors.New("Custom.Error4")},
	ErrorType5: customError{Error: xerrors.New("Custom.Error5")},
}
