package main

import (
	"fmt"
	"os"
)

func main() {
	// pointer variables
	var x, y, z = 23, 43.2, "Hay !"
	var p = &z
	fmt.Println(*p, x, y)
	*p = " Buddy"
	fmt.Println("value of p", *p, p, &p)
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		// read file
	}
	f.Close()
	var xP = &x
	var pointer = localPointer(&x)
	fmt.Println(pointer == xP)
	fmt.Println(pointerGen() == pointerGen())
	fmt.Println(structGen() == structGen())

}
func localPointer(x *int) *int {
	return x
}
func pointerGen() *int {
	return new(int)
}
func structGen() *struct{} {
	return new(struct{})
}
