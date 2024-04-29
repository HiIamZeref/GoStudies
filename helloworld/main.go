package main // package: main is a special package name in Go. It is used to create an executable program.

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

// How to run the program?
// R: go run main.go


// Some important commands:
// go build main.go -- to build the executable file
// ./main -- to run the executable file after building
// go fmt main.go -- to format the code in the file
// go install main.go -- to install the executable file in the bin directory
// go get github.com/... -- to get the package from the github
// go test main.go -- to test the code in the file


// What does "package" mean? (Package == Project == Workspace)
// R: A package is a collection of Go source files that reside in the same directory. It is used to organize the code and to make it reusable.
// Types of packages:
// 1. Executable package: It is used to create an executable program. It contains the main function. (package main)
// 2. Reusable package: It is used to create a reusable code. It does not contain the main function. (package <package_name>)


// What does "import" mean?
// R: The import keyword is used to import the packages in the Go program. It is used to use the code from the other packages.


// File organization in Go:
// Top level: package declaration
// Next level: import statements
// Next level: function declarations

