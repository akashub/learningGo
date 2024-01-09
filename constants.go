package main

import (
	"fmt"
)

func main() {
	const (
		name = "John"
		age = 50
		country = "Canada"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	// Constants, as the name indicate, cannot be reassigned again to any other value. In the program below, we are trying to assign another value 89 to a. This is not allowed since a is a constant. This program will fail to run with compilation error cannot assign to a.
	//const a = 55 -> assigned
	//a = 89 -> reassignment not allowed

	//The value of a constant should be known at compile time. Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.

	//var a = math.Sqrt(4) is allowed
	//but const b = math.Sqrt(4) is not allowed


	//Any value enclosed between double quotes is a string constant in Go. For example, strings like "Hello World", "Sam" are all constants in Go.

	// string constant belongs to untyped type
	// A string constant like "Hello world" does not have any type

	//untyped constants have a default type associated with them and they supply it if and only if a line of code demands it. In the statement var name = n, name needs a type and it gets it from the default type of the string constant n which is a string. -> Check the photo AssigningVariable...

	//const typedhello string = "Hello World" -> constant of type string different from
	//const untypedhello = "Hello World" -> this is of untyped type

	var defaultName = "Aakash" //allowed
	type myString string
 	var customName myString = "Sam" //allowed -> myString is an alias for string type
	//customName = defaultName -> not allowed

	//Now we have a variable defaultName of type string and another variable customName of type myString. Even though we know that myString is an alias of string, Go’s strong typing policy disallows variables of one type to be assigned to another. Hence the assignment customName = defaultName is not allowed and the compiler throws the error ./prog.go:7:20: cannot use defaultName (type string) as type myString in assignment

	const trueConst = true
	type myBool bool

}
