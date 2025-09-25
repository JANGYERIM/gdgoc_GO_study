package main

import (
	"fmt"
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>

void printHelloWorld() {
   printf("Hello, World!\n");
}

void printString(char* str) {
   printf("%s\n", str);
}

char *getName(int idx) {
   if (idx == 1)
      return "Baheer Kamal";
   if (idx == 2)
      return "Tanmay Bakshi";
   return "Unknown";
}
*/
import "C"

func main() {
	C.printHelloWorld()
	a := C.CString("This is from Golang")
	C.printString(a)
	C.free(unsafe.Pointer(a)) // 메모리 할당 해제

	cstr := C.getName(C.int(2))
	fmt.Println(C.GoString(cstr)) // C 문자열을 Go 문자열로 변환
}
