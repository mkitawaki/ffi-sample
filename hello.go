package main

/*
#include <stdio.h>
#include <stdlib.h>

struct int_word {
	int cn;
	char *word;
};
void hello(struct int_word *iw)
{
    printf("Hello, cgo : %s %d\n",  iw -> word, iw -> cn);
}
*/
import "C"
import "unsafe"

type IntWord struct {
	cn   int32
	word string
}

func main() {
	vw := &IntWord{1, "World"}
	p := unsafe.Pointer(vw)
        C.hello((*C.struct_int_word)(p))

}
