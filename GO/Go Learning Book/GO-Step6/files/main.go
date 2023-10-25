package main

import (
	"fmt"
	"io"

	//"io/ioutil"
	"os"
)

func readFile(filename string) {
	databyte, err := os.ReadFile(filename) //read the file //os.ReadFile() function reads the entire contents of a file as a byte slice. It returns the byte slice and an error, if any.
	checkErr(err)
	fmt.Println("Text dat In file: ", string(databyte)) //print the contents of the file
}
func checkErr(err error) {
	if err != nil {
		panic(err) //error handling for file creation failure //panic is a built in function which stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.
	}
}

func main() {
	fmt.Println("Welcome to files Study!")

	content := "Hello from Shykat - Golang Developer"

	file, err := os.Create("./mygofile.txt") //create a file in the current directory //os.Create() function creates a file if it doesn't exist, or truncates the file if it exists. //os.Create() function returns a file descriptor that can be used to read and write to the file.
	checkErr(err)
	length, err := io.WriteString(file, content) //write content to file //io.WriteString() function writes the contents of a string to a file. It returns the number of bytes written and an error, if any.
	checkErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close() //close the file when done //file.Close() function closes a file. It returns an error if it fails.

	readFile("./mygofile.txt") //read the file we just created and print the contents to the console

}
