package main

/*
#cgo CFLAGS: -I${SRCDIR}/cpp
#cgo LDFLAGS: -L${SRCDIR}/cpp -lhello -lstdc++
#include <stdio.h>
#include <stdlib.h>
#include "hello.h"
*/
import "C"
import (
	"fmt"
)

func main() {
	p := C.Create()
	name := C.GoString(C.GetName(p))
	age := C.GetAge(p)
	C.Destroy(p)
	fmt.Printf("Name is %s, age is %d\n", name, age)
}
