package main

//#cgo CFLAGS: -I./cefapi
//#cgo LDFLAGS: -L${SRCDIR}/cefapi -L${SRCDIR}/Release -lcefapi -lcef
/*
#include <stdio.h>
#include "cefapi.h"
*/
import "C"

//注意此处C需要和其他的包分开
import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	fmt.Println(C.number_add_mod(10, 5, 12))
	args := os.Args
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	C.startCef(argc, (**C.char)(unsafe.Pointer(&argv[0])))
}
