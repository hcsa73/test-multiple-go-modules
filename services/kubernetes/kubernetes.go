package kubernetes

import (
	"fmt"

	core "github.com/hcsa73/test-multiple-go-modules/core"
)

const DUMMY_CONST = "KUBERNETES-V0.1.1"

func GetDummyConst() string {
	return DUMMY_CONST
}

func Info() {
	fmt.Println("Hello from kubernetes!")
	fmt.Println(GetDummyConst())
	fmt.Printf("Core const: %v\n", core.GetDummyConst())
}
