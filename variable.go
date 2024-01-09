package main

import (
	"fmt"
	"math"
)

func main() {
	// var age int // variable declaration
	// fmt.Println("My age is", age)
	// here no value is given to age and hence go initialises to 0
	var age int // variable declaration
	fmt.Println("My age is", age)
	age = 29 //assignment
	fmt.Println("My age is", age)
	age = 54 //assignment
	fmt.Println("My new age is", age)
	var width, height = 100, 50
	fmt.Println(width*height)
	var (
			name = "Aakash"
			sex = "M"
	)
	height = 200
	fmt.Println(name, sex, "height in cm is", height)
	active := "High"
	fmt.Println("activity level", active)
	// name, age, sex := "Nidhi", 50, "F"
	// The above statement cannot work since there are no new variables declared in the left of concise variable declaration :=, atleast 1 should be a new variable
	x, y, z := 1.12, 2.00, 3.09
	fmt.Println(x, y, z) 

	c := math.Min(x, y) //can only process float64
	fmt.Println(c)
}

//if a variable is declared then it has to be used.
