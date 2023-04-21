package core

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const DUMMY_CONST = "CORE-V0.1.2"

func GetDummyConst() string {
	return DUMMY_CONST
}

func GetDummyConstReversed() string {
	return stringutil.Reverse(DUMMY_CONST)
}

func Info() {
	fmt.Println("Hello from core!")
	fmt.Println(GetDummyConstReversed())
}
