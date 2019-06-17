package main

import (
	"log"
	"reflect"
)

type User struct {
	ID   int
	Name string
}

func main() {
	var u1 interface{}
	u2 := Fn()
	var u3 *User = nil

	log.Printf("%+v\n", reflect.TypeOf(u1))
	log.Printf("%+v\n", reflect.TypeOf(u2))
	log.Printf("%+v\n", reflect.TypeOf(u3))
	log.Printf("%+v\n", reflect.TypeOf(u4))
	log.Printf("%+v\n", u1)
	log.Printf("%+v\n", u2)
}

func Fn() *User {
	return nil
}
