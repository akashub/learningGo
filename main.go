package main //this indicates that this file is supposed to be an executable program rather than a reusable library or module, there also needs to be a main function within the file that acts as an entry point for execution

import "fmt" 
///This line imports the "fmt" package, which stands for "format." The "fmt" package provides functions for formatting input and output.

func main() {
	fmt.Println("Hello World")
}

//It’s a convention in Go to name the file that contains the main function as main.go, but other names work as well.

// can be run using ~/go/bin/learningGo
// also by go build, learningGo.exe
// go run main.go -> Works similar to go build/go install, but it requires extension .go 

//Under the hood, go run works much similar to go build. Instead of compiling and installing the program to the current directory, it compiles the file to a temporary location and runs the file from that location. If you are interested to know the location where go run compiles the file to, please run go run with the --work argument.
//Gives out 
//WORK=C:\Users\aakas\AppData\Local\Temp\go-build1893233669
//Hello World