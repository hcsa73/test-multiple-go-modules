package main

import (
	"fmt"

	"github.com/hcsa73/test-multiple-go-modules/core"
	"github.com/hcsa73/test-multiple-go-modules/services/kubernetes"
)

func main() {
	fmt.Println("Start main")
	fmt.Println("==== core info ====")
	core.Info()
	fmt.Println("==== kubernetes info ====")
	kubernetes.Info()
}
