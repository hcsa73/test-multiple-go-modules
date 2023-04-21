package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const DUMMY_CONST = "CORE"

func GetDummyConst() string {
	return DUMMY_CONST
}

func GetDummyConstReversed() string {
	return stringutil.Reverse(DUMMY_CONST)
}

func main() {
	fmt.Println("Hello from core!")
	fmt.Println(GetDummyConstReversed())
}
