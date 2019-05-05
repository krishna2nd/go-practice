package main

/*
#include <stdio.h>
static void p(char *s) {
	printf("%s", s);
}
*/
import "C"
import (
	"fmt"
)


func main() {
	cs := C.CString("Hello from stdio")
	C.p(cs)
	v:= "test";
	fmt.Println(v);
}
