package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"

func main() {
	println(C.random())
}
