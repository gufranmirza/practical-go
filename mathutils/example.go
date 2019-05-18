package main

import (
	"fmt"
	"github.com/gufranmirza/practical-go/mathutils"
)


func main() {
	fmt.Println("Addition of two numbers:", mathutils.Addition(10, 10))
	fmt.Println("Substraction of two numbers:", mathutils.Substraction(100, 40))
	fmt.Println("Multiplication of two numbers:", mathutils.Multiplication(10, 20))
	fmt.Println("Square of a number:", mathutils.Square(10))
}
