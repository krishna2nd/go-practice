package main

import (
	"fmt"
	"seccond"
	"third"
)

func init() {
	fmt.Print("init")
}
func main() {
	seccond.Test()
	third.Test()
	fmt.Println("test")
}
