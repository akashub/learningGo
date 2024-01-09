// The following are the basic types available in Go

// bool
// Numeric Types
// int8, int16, int32, int64, int
// uint8, uint16, uint32, uint64, uint
// float32, float64
// complex64, complex128
// byte
// rune
// string

package main

import (
	"fmt"
	"unsafe"
)


func main() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
	var x int = 89
	y := 99
	fmt.Println("value of x is", x, "value of y is", y)
	// x is defined to have type int and y's int type is inferred from assigned value
	fmt.Printf("type of x is '%T', size of x is %d", x, unsafe.Sizeof(x))
	fmt.Printf("\ntype of y is %T, size of y is %d", x, unsafe.Sizeof(x))

	// 	%T: This specifier is used to format the operand's type.
	// %d: This specifier is used to format the operand as a decimal integer.
	// Printf statement can be used to use specifiers
	// Go has a package unsafe which has a Sizeof function which returns in bytes the size of the variable passed to it. unsafe package should be used with care as the code using it might have portability issues

	f, g := 5.67, 8.97
	fmt.Printf("type of f %T g %T\n", f, g)
	sum := f + g
	diff := f - g
	fmt.Println("sum", sum, "diff", diff)
	// the output of diff would be -3.3000000000000007 rather than -3.30
	// the result -3.3000000000000007 is a result of the limited precision of floating-point numbers. When you perform operations like addition or subtraction on floating-point numbers, the result may not always be exact due to the inherent limitations of representing real numbers in binary.

	//use the below line of code for diff of float
	fmt.Println("sum", sum, "diff", fmt.Sprintf("%.2f", diff)) 
	//rounds off to -3.30 because of .2f
	//Sprintf is used to store the formatted string to a value and then print it unlike Printf that directly prints it, using Sprintf you can store values to a variable that'l get printed when you call them


	//complex64: complex numbers which have float32 real and imaginary parts
	//complex128: complex numbers with float64 real and imaginary parts

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	// byte is an alias of uint8
	// rune is an alias of int32 -> more about it in strings.go

	i := 55 //int
	j := 67.8//float64
	//sum := i + j is not allowed becase go is very strict about explicit typing, no automatic type conversion/promotion
	//type conversion is required
	sum1 := i + int(j)
	fmt.Println(sum1)

}
