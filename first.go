package main

import (
	"fmt"
)

// var x,y,z int
// var c,python, java bool





func main() {

var num int = 1

var numPtr *int = &num

fmt.Println(numPtr)
fmt.Println(*numPtr)
fmt.Println(&num) 
}
