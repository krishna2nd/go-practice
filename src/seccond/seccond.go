package seccond

import (
	"fmt"
	"third"
)

func init() {
	fmt.Print("seccond")
}

func Test() {
	fmt.Print("Test..")
	third.Test()
}
