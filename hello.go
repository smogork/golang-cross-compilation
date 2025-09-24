package main

import "fmt"

/*
#include <stdlib.h>
*/
import "C"

func main() {
	C.srand(C.uint(4))
	var num = C.rand()

	fmt.Printf("Hello, World! Your number is %d\n", num)
}
