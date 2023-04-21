package core

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const DUMMY_CONST = "CORE-V1.0.0"

func GetDummyConstReversed() string {
	return stringutil.Reverse(DUMMY_CONST)
}

func Info() {
	fmt.Println("Hello from core!")
	fmt.Println(GetDummyConstReversed())
}
