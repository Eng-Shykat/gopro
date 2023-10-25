package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps Study!")
	//create a map with key as string and value as string and store in languages variable

	languages := make(map [string] string) //create a map with key as string and value as string and store in languages variable
	languages["PY"] = "Python"
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"
	languages["CS"] = "C#"

	fmt.Println("List of all languages:",languages)
	fmt.Println("JS stands for",languages["JS"]) //print the value of JS key from languages map variable 

	delete(languages,"CS") //delete the key CS from languages map variable 
	fmt.Println("List of all languages:",languages) 
	fmt.Print("Done\n\n")

	//loops through a map
	for key, value := range languages { //loop through languages map variable 
		fmt.Printf("Key: %v, Value: %v\n", key, value) //print the key and value of each element of languages map variable
	}

}