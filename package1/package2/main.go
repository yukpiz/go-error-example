package main

import (
	"log"

	"github.com/yukpiz/go-error-example/package1"
)

func main() {
	u1 := package1.Package1User{}
	log.Printf("%+v\n", u1)

	u2 := package1.internalUser{}
	log.Printf("%+v\n", u2)
}
