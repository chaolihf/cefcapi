package main

//#cgo CFLAGS: -I./cefapi
//#cgo LDFLAGS: -L${SRCDIR}/cefapi -L${SRCDIR}/Release -lcefapi -lcef
//
//#include "cefapi.h"
import "C"

//注意此处C需要和其他的包分开
import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(C.number_add_mod(10, 5, 12))
}
