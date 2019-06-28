package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int g_var = 0;

int add(int a, int b)
{
	printf("[%s] g_var = %d\n", __FUNCTION__, ++g_var);
	return a + b;
}

char* hello(char *str)
{
	printf("[%s] g_var = %d\n", __FUNCTION__, ++g_var);
	return strcat(str, " World\n");
}
*/
import "C"

import "fmt"

func main() {
	fmt.Println("golang g_var = ", C.g_var)

	a := C.int(1)
	b := C.int(2)
	fmt.Println(a, "+", b, "=", C.add(a, b))

	fmt.Println("golng g_var = ", C.g_var)

	s := C.CString("Hello")
	cString := C.hello(s)
	gostr := C.GoString(cString)
	fmt.Println(gostr)
}