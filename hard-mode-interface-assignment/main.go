package main

import (
	"io"
	"os"
)

func main() {
	// Saving my args
	args := os.Args
	
	// Getting my filename
	file_name := args[1] 

	// Opening file
	file, err := os.Open(file_name)

	// Checking for errors
	if err != nil {
		panic(err) // If there is an error, panic -> stop the program
	}

	// Using 'io.Copy' to copy the file to the standard output
	io.Copy(os.Stdout, file)
}